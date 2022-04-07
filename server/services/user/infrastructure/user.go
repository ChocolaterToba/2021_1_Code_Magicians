package repository

import (
	"context"
	"pinterest/services/user/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user domain.User, passwordHash []byte) (userID uint64, err error)
	GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error)
	UpdateUser(ctx context.Context, user domain.User) (err error)
}

type UserRepo struct {
	postgresDB *pgxpool.Pool
}

func NewUserRepo(postgresDB *pgxpool.Pool) *UserRepo {
	return &UserRepo{postgresDB: postgresDB}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user domain.User, passwordHash []byte) (userID uint64, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	createUserQuery := `INSERT INTO users (username, password_hash, email, first_name, last_name)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING id`

	row := tx.QueryRow(context.Background(), createUserQuery, user.Username, passwordHash, user.Email, user.FirstName, user.LastName)
	err = row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, domain.TransactionCommitError
	}
	return userID, nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, user domain.User) (err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	updateUserQuery := `UPDATE users
						SET email = $2, first_name = $3, last_name = $4
						WHERE id = $1`

	result, err := tx.Exec(context.Background(), updateUserQuery, user.UserID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return domain.UserNotFoundError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.TransactionCommitError
	}
	return nil
}

func (repo *UserRepo) GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.User{}, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	getUserByID := `SELECT id, username, email, first_name, last_name
					FROM users
					WHERE id = $1`

	row := tx.QueryRow(context.Background(), getUserByID, userID)
	err = row.Scan(&user.UserID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, domain.UserNotFoundError
		}

		return domain.User{}, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.User{}, domain.TransactionCommitError
	}
	return user, nil
}
