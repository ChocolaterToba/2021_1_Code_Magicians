package application

import (
	"context"
	"crypto/rand"
	"pinterest/services/auth/domain"
	repository "pinterest/services/auth/infrastructure"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthAppInterface interface {
	LoginUser(ctx context.Context, username string, password string) (cookie domain.CookieInfo, err error)
	SearchCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error)
	SearchCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error)
	LogoutUser(ctx context.Context, cookieValue string) (err error)
	ChangeCredentials(ctx context.Context, userID uint64, username string, password string) (err error)
}

type AuthApp struct {
	repo repository.AuthRepoInterface
}

func NewAuthApp(repo repository.AuthRepoInterface) *AuthApp {
	return &AuthApp{
		repo: repo,
	}
}

func (app *AuthApp) LoginUser(ctx context.Context, username string, password string) (cookie domain.CookieInfo, err error) {
	userId, passwordHash, err := app.repo.GetPasswordHash(ctx, username)
	if err != nil {
		return domain.CookieInfo{}, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return domain.CookieInfo{}, domain.IncorrectPasswordError
		}

		return domain.CookieInfo{}, err
	}

	cookie.UserID = userId
	cookie.Cookie.Value = randString(domain.CookieSessionLength)
	cookie.Cookie.Expires = time.Now().Add(time.Hour * domain.CookieExpiryHours)

	err = app.repo.AddCookieInfo(ctx, cookie)
	if err != nil {
		return domain.CookieInfo{}, err
	}

	return cookie, nil
}

func randString(n int) string {
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func (app *AuthApp) SearchCookieByValue(ctx context.Context, cookieValue string) (cookie domain.CookieInfo, err error) {
	// TODO: add expiry check
	return app.repo.GetCookieByValue(ctx, cookieValue)
}

func (app *AuthApp) SearchCookieByUserID(ctx context.Context, userID uint64) (cookie domain.CookieInfo, err error) {
	// TODO: add expiry check
	return app.repo.GetCookieByUserID(ctx, userID)
}

func (app *AuthApp) LogoutUser(ctx context.Context, cookieValue string) (err error) {
	return app.repo.DeleteCookie(ctx, cookieValue)
}

func (app *AuthApp) ChangeCredentials(ctx context.Context, userID uint64, username string, password string) (err error) {
	//TODO: add transactions here?
	dbUsername, dbPasswordHash, err := app.repo.GetUserCredentialsByID(ctx, userID)
	if err != nil {
		return err
	}

	if username != "" {
		dbUsername = username
	}
	if password != "" {
		dbPasswordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
	}

	return app.repo.UpdateUserCredentials(ctx, userID, dbUsername, dbPasswordHash)
}
