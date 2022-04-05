package repository

import (
	"context"
	"pinterest/services/user/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user domain.User, passwordHash []byte) (userID uint64, err error)
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
