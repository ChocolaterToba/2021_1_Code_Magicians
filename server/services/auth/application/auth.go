package application

import (
	"context"
	"pinterest/services/auth/domain"
	repository "pinterest/services/auth/infrastructure"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthAppInterface interface {
	CheckUserCredentials(ctx context.Context, username string, password string) (userID uint64, err error)
	AddCookieInfo(ctx context.Context, cookieInfo domain.CookieInfo) error
	SearchCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error)
	SearchCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error)
	RemoveCookie(ctx context.Context, cookieInfo domain.CookieInfo) error
	GetUserIDByVkID(ctx context.Context, vkID uint64) (userID uint64, err error)
	AddVkID(ctx context.Context, userID uint64, vkID uint64) error
}

type AuthApp struct {
	postgresDB *pgxpool.Pool
	repo       repository.AuthRepoInterface
}

func NewService(postgresDB *pgxpool.Pool, repo repository.AuthRepoInterface) *AuthApp {
	return &AuthApp{
		postgresDB: postgresDB,
		repo:       repo,
	}
}

func (app *AuthApp) CheckUserCredentials(ctx context.Context, username string, password string) (userID uint64, err error) {
	return app.repo.CheckUserCredentials(ctx, username, password)
}

func (app *AuthApp) AddCookieInfo(ctx context.Context, cookieInfo domain.CookieInfo) error {
	return app.repo.AddCookieInfo(ctx, cookieInfo)
}

func (app *AuthApp) SearchCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error) {
	return app.repo.GetCookieByValue(ctx, cookieValue)
}

func (app *AuthApp) SearchCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error) {
	return app.repo.GetCookieByUserID(ctx, userID)
}

func (app *AuthApp) RemoveCookie(ctx context.Context, cookieInfo domain.CookieInfo) error {
	return app.repo.DeleteCookie(ctx, cookieInfo.UserID)
}

func (app *AuthApp) GetUserIDByVkID(ctx context.Context, vkID uint64) (userID uint64, err error) {
	return app.repo.GetUserIDByVkID(ctx, vkID)
}

func (app *AuthApp) AddVkID(ctx context.Context, userID uint64, vkID uint64) error {
	return app.repo.AddVkID(ctx, userID, vkID)
}
