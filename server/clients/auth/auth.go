package auth

import (
	"context"
	"net/http"
	"pinterest/domain"
	authdomain "pinterest/services/auth/domain"
	authproto "pinterest/services/auth/proto"
	"strings"

	"github.com/pkg/errors"
)

type AuthClientInterface interface {
	LoginUser(ctx context.Context, username string, password string) (cookie *domain.CookieInfo, err error)
	SearchCookieByValue(ctx context.Context, cookieValue string) (cookie *domain.CookieInfo, err error)
	SearchCookieByUserID(ctx context.Context, userID uint64) (cookie *domain.CookieInfo, err error)
	LogoutUser(ctx context.Context, cookieValue string) error
	ChangeCredentials(ctx context.Context, userID uint64, username, password string) (err error)
}

type AuthClient struct {
	authClient authproto.AuthClient
}

func NewAuthClient(authClient authproto.AuthClient) *AuthClient {
	return &AuthClient{
		authClient: authClient,
	}
}

func (client *AuthClient) LoginUser(ctx context.Context, username string, password string) (cookie *domain.CookieInfo, err error) {
	pbCookie, err := client.authClient.LoginUser(context.Background(),
		&authproto.UserAuth{Username: username, Password: password})

	if err != nil {
		if strings.Contains(err.Error(), authdomain.IncorrectPasswordError.Error()) {
			return nil, domain.ErrIncorrectPassword
		}
		return nil, errors.Wrap(err, "auth client error: ")
	}

	cookie = domain.ToCookieInfo(pbCookie, true, true, http.SameSiteDefaultMode) // TODO: move settings to constants
	return cookie, nil
}

func (client *AuthClient) SearchCookieByValue(ctx context.Context, cookieValue string) (cookie *domain.CookieInfo, err error) {
	pbCookie, err := client.authClient.SearchCookieByValue(context.Background(),
		&authproto.CookieValue{CookieValue: cookieValue})

	if err != nil {
		if strings.Contains(err.Error(), authdomain.CookieNotFoundError.Error()) {
			return nil, domain.ErrCookieNotFound
		}
		return nil, errors.Wrap(err, "auth client error: ")
	}

	cookie = domain.ToCookieInfo(pbCookie, true, true, http.SameSiteDefaultMode) // TODO: move settings to constants
	return cookie, nil
}

func (client *AuthClient) SearchCookieByUserID(ctx context.Context, userID uint64) (cookie *domain.CookieInfo, err error) {
	pbCookie, err := client.authClient.SearchCookieByUserID(context.Background(),
		&authproto.UserID{Uid: userID})

	if err != nil {
		if strings.Contains(err.Error(), authdomain.CookieNotFoundError.Error()) {
			return nil, domain.ErrCookieNotFound
		}
		return nil, errors.Wrap(err, "auth client error: ")
	}

	cookie = domain.ToCookieInfo(pbCookie, true, true, http.SameSiteDefaultMode) // TODO: move settings to constants
	return cookie, nil
}

func (client *AuthClient) LogoutUser(ctx context.Context, cookieValue string) error {
	_, err := client.authClient.LogoutUser(context.Background(),
		&authproto.CookieValue{CookieValue: cookieValue})

	if err != nil {
		if strings.Contains(err.Error(), authdomain.CookieNotFoundError.Error()) {
			return domain.ErrCookieNotFound
		}
		return errors.Wrap(err, "auth client error: ")
	}

	return nil
}

func (client *AuthClient) ChangeCredentials(ctx context.Context, userID uint64, username, password string) (err error) {
	_, err = client.authClient.ChangeCredentials(context.Background(),
		&authproto.Credentials{
			UserID:   userID,
			Username: username,
			Password: password,
		})

	if err != nil {
		return errors.Wrap(err, "auth client error: ")
	}

	return nil
}
