package profile

import (
	"context"
	"encoding/json"
	"net/http"
	authclient "pinterest/clients/auth"
	userclient "pinterest/clients/user"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
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

// EditUser changes users's non-credential (like email, first name, etc) data to one specified
func (facade *ProfileFacade) EditUser(w http.ResponseWriter, r *http.Request) {
	userInput := new(domain.User)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userCookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)
	userInput.UserID = userCookie.UserID
	userInput.Username = ""
	userInput.Password = ""

	err = facade.userClient.EditUser(context.Background(), *userInput)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrUserNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

// GetUserByID recieves user data from user service. E-mail gets hidden for personal data protection
func (facade *ProfileFacade) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr, passedID := vars[string(domain.IDKey)]

	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	user, err := facade.userClient.GetUserByID(context.Background(), userID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		if err == domain.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Email = ""

	responseBody, err := json.Marshal(user)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	return
}

// GetUserByUsername recieves user data from user service. E-mail gets hidden for personal data protection
func (facade *ProfileFacade) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, passedUsername := vars[string(domain.UsernameKey)]

	if !passedUsername {
		facade.logger.Info("Could not get username from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := facade.userClient.GetUserByUsername(context.Background(), username)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		if err == domain.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Email = ""

	responseBody, err := json.Marshal(user)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	return
}

// GetCurrentUser recieves current user data from user service. Sensitive data is preserved in return values
func (facade *ProfileFacade) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	user, err := facade.userClient.GetUserByID(context.Background(), cookie.UserID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		if err == domain.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(user)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	return
}

const maxPostAvatarBodySize = 8 * 1024 * 1024 // 8 mB
// HandlePostAvatar takes avatar from request and assigns it to current user
func (facade *ProfileFacade) HandlePostAvatar(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	bodySize := r.ContentLength

	if bodySize <= 0 { // No avatar was passed
		facade.logger.Info(domain.ErrNoFile.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if bodySize > int64(maxPostAvatarBodySize) { // Avatar is too large
		facade.logger.Info(domain.ErrFileTooLarge.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(bodySize)
	file, header, err := r.FormFile("avatarImage")
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = facade.userClient.UpdateAvatar(context.Background(), cookie.UserID, header.Filename, file)
	if err != nil {
		// TODO: parse wrong extension errors
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
