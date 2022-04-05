package profile

import (
	"context"
	"encoding/json"
	"net/http"
	authclient "pinterest/clients/auth"
	userclient "pinterest/clients/user"
	"pinterest/domain"

	"go.uber.org/zap"
)

// ProfileFacade calls user app
type ProfileFacade struct {
	userClient userclient.UserClientInterface
	authClient authclient.AuthClientInterface
	logger     *zap.Logger
}

func NewProfileFacade(userClient userclient.UserClientInterface, authClient authclient.AuthClientInterface, logger *zap.Logger) *ProfileFacade {
	return &ProfileFacade{
		userClient: userClient,
		authClient: authClient,
		logger:     logger,
	}
}

// Create user creates user using provided data, also logs user in
func (facade *ProfileFacade) CreateUser(w http.ResponseWriter, r *http.Request) {
	userInput := new(domain.User)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := facade.userClient.CreateUser(context.Background(), *userInput)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError) // TODO: recognize duplicate usernames etc
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

	userOutput := domain.UserIDResponse{UserID: userID}
	responseBody, err := json.Marshal(userOutput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookieInfo.Cookie)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
