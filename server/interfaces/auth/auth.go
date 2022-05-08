package auth

import (
	"context"
	"encoding/json"
	"net/http"
	authclient "pinterest/clients/auth"
	"pinterest/domain"
	"pinterest/interfaces/middleware"

	"time"

	"go.uber.org/zap"
)

// AuthFacade calls auth app
type AuthFacade struct {
	authClient authclient.AuthClientInterface
	logger     *zap.Logger
}

func NewAuthFacade(authClient authclient.AuthClientInterface, logger *zap.Logger) *AuthFacade {
	return &AuthFacade{
		authClient: authClient,
		logger:     logger,
	}
}

// LoginUser logs user in using provided username and password
func (facade *AuthFacade) LoginUser(w http.ResponseWriter, r *http.Request) {
	userInput := new(domain.UserCredentialsInput)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookieInfo, err := facade.authClient.LoginUser(context.Background(), userInput.Username, userInput.Password)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrIncorrectPassword:
			w.WriteHeader(http.StatusUnauthorized)
		case domain.ErrUserNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, cookieInfo.Cookie)
	w.WriteHeader(http.StatusNoContent)
}

// LogoutUser logs current user out of their session
func (facade *AuthFacade) LogoutUser(w http.ResponseWriter, r *http.Request) {
	userCookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	err := facade.authClient.LogoutUser(context.Background(), userCookie.Cookie.Value)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI),
			zap.Uint64("for user", userCookie.UserID),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCookie.Cookie.Expires = time.Now().AddDate(0, 0, -1) // Making cookie expire
	http.SetCookie(w, userCookie.Cookie)

	w.WriteHeader(http.StatusNoContent)
}

// CheckUser checks if current user is logged in
func (facade *AuthFacade) CheckUser(w http.ResponseWriter, r *http.Request) {
	_, found := middleware.CheckCookies(r, facade.authClient)
	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ChangeCredentials changes user's username and/or password
func (facade *AuthFacade) ChangeCredentials(w http.ResponseWriter, r *http.Request) {
	userInput := new(domain.UserCredentialsInput)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userCookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)
	err = facade.authClient.ChangeCredentials(context.Background(), userCookie.UserID, userInput.Username, userInput.Password)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
