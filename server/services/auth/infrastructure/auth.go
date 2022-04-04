package repository

import (
	"context"
	"pinterest/services/auth/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepoInterface interface {
	CheckUserCredentials(ctx context.Context, username string, password string) (userID uint64, err error)
	AddCookieInfo(ctx context.Context, cookieInfo domain.CookieInfo) error
	GetCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error)
	GetCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error)
	DeleteCookie(ctx context.Context, cookieValue string) error
}

type AuthRepo struct {
	postgresDB *pgxpool.Pool
}

func NewAuthRepo(postgresDB *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{postgresDB: postgresDB}
}

func (repo *AuthRepo) CheckUserCredentials(ctx context.Context, username string, password string) (userID uint64, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return 0, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	getUserPasswordQuery := `SELECT user_id, password_hash
							 FROM Users 
							 WHERE username=$1`

	passwordHash := make([]byte, 0)

	row := tx.QueryRow(context.Background(), getUserPasswordQuery, username)
	err = row.Scan(&userID, &passwordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, domain.UserNotFoundError
		}

		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return 0, domain.IncorrectPasswordError
		}

		return 0, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, domain.TransactionCommitError
	}
	return userID, nil
}

func (repo *AuthRepo) AddCookieInfo(ctx context.Context, cookieInfo domain.CookieInfo) error {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	addCookieQuery := `UPDATE user
					   SET cookie_value = $2, cookie_expiry = $3
					   WHERE user_id = $1`

	result, err := tx.Exec(context.Background(), addCookieQuery, cookieInfo.UserID, cookieInfo.Cookie.Value, cookieInfo.Cookie.Expires)
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

func (repo *AuthRepo) GetCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.CookieInfo{}, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	getCookieByValueQuery := `SELECT user_id, cookie_expiry
							  FROM user
							  WHERE cookie_value = $1`

	row := tx.QueryRow(context.Background(), getCookieByValueQuery, cookieValue)
	err = row.Scan(&cookie.UserID, &cookie.Cookie.Expires)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.CookieInfo{}, domain.CookieNotFoundError
		}

		return domain.CookieInfo{}, err
	}

	cookie.Cookie.Value = cookieValue

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.CookieInfo{}, domain.TransactionCommitError
	}
	return cookie, nil
}

func (repo *AuthRepo) GetCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.CookieInfo{}, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	getCookieByUserIDQuery := `SELECT cookie_value, cookie_expiry
							   FROM user
							   WHERE cookie_value = $1`

	row := tx.QueryRow(context.Background(), getCookieByUserIDQuery, userID)
	err = row.Scan(&cookie.Cookie.Value, &cookie.Cookie.Expires)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.CookieInfo{}, domain.CookieNotFoundError
		}

		return domain.CookieInfo{}, err
	}

	cookie.UserID = userID

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.CookieInfo{}, domain.TransactionCommitError
	}
	return cookie, nil
}

func (repo *AuthRepo) DeleteCookie(ctx context.Context, cookieValue string) error {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	deleteCookieQuery := `UPDATE user
						  SET cookie_value = '', cookie_expiry = now()
						  WHERE user.cookie_value = $1`

	result, err := tx.Exec(context.Background(), deleteCookieQuery, cookieValue)
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
