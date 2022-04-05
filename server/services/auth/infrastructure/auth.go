package repository

import (
	"context"
	"pinterest/services/auth/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepoInterface interface {
	GetPasswordHash(ctx context.Context, username string) (userID uint64, passwordHash []byte, err error)
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

func (repo *AuthRepo) GetPasswordHash(ctx context.Context, username string) (userID uint64, passwordHash []byte, err error) {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return 0, nil, domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	getUserPasswordQuery := `SELECT id, password_hash
							 FROM users 
							 WHERE username=$1`

	passwordHash = make([]byte, 0)

	row := tx.QueryRow(context.Background(), getUserPasswordQuery, username)
	err = row.Scan(&userID, &passwordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil, domain.UserNotFoundError
		}

		return 0, nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, nil, domain.TransactionCommitError
	}
	return userID, passwordHash, nil
}

func (repo *AuthRepo) AddCookieInfo(ctx context.Context, cookieInfo domain.CookieInfo) error {
	tx, err := repo.postgresDB.Begin(context.Background())
	if err != nil {
		return domain.TransactionBeginError
	}
	defer tx.Rollback(context.Background())

	addCookieQuery := `UPDATE users
					   SET cookie_value = $2, cookie_expiry = $3
					   WHERE id = $1`

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

	getCookieByValueQuery := `SELECT id, cookie_expiry
							  FROM users
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
							   FROM users
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

	deleteCookieQuery := `UPDATE users
						  SET cookie_value = '', cookie_expiry = now()
						  WHERE users.cookie_value = $1`

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
