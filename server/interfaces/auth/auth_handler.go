package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"pinterest/usage"
	"pinterest/domain/entity"
	"pinterest/interfaces/middleware"
	"time"

	"go.uber.org/zap"
)

// AuthInfo keep information about apps and cookies needed for auth package
type AuthInfo struct {
	authApp      usage.AuthAppInterface
	userApp      usage.UserAppInterface
	cookieApp    usage.CookieAppInterface
	s3App        usage.S3AppInterface
	boardApp     usage.BoardAppInterface     // For initial user's board
	websocketApp usage.WebsocketAppInterface // For setting CSRF token during  login
	logger       *zap.Logger
}

func NewAuthInfo(userApp usage.UserAppInterface, cookieApp usage.CookieAppInterface,
	s3App usage.S3AppInterface, boardApp usage.BoardAppInterface,
	websocketApp usage.WebsocketAppInterface, logger *zap.Logger) *AuthInfo {
	return &AuthInfo{
		userApp:      userApp,
		cookieApp:    cookieApp,
		s3App:        s3App,
		boardApp:     boardApp,
		websocketApp: websocketApp,
		logger:       logger,
	}
}

// HandleCreateUser creates user with parameters passed in JSON
func (info *AuthInfo) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	userInput := new(entity.UserRegInput)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, _ := userInput.Validate()
	if !valid {
		info.logger.Info(entity.ValidationError.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newUser entity.User
	err = newUser.UpdateFrom(userInput)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, _ := info.userApp.GetUserByUsername(newUser.Username)
	if user != nil {
		info.logger.Info(entity.UsernameEmailDuplicateError.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusConflict)
		return
	}

	cookie, err := info.cookieApp.GenerateCookie()
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newUser.UserID, err = info.userApp.CreateUser(&newUser)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		if err.Error() == entity.UsernameEmailDuplicateError.Error() {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = info.cookieApp.AddCookieInfo(&entity.CookieInfo{newUser.UserID, cookie})
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	// Replacing token in websocket connection info
	token := r.Header.Get("X-CSRF-Token")
	err = info.websocketApp.ChangeToken(newUser.UserID, token)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleLoginUser logs user in using provided username and password
func (info *AuthInfo) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	userInput := new(entity.UserLoginInput)
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
    log.Println("LOG1")
	user, err := info.authApp.LoginUser(userInput.Username, userInput.Password)

	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err.Error() {
		case entity.IncorrectPasswordError.Error():
			w.WriteHeader(http.StatusUnauthorized)
		case string(entity.UserNotFoundError):
			w.WriteHeader(http.StatusNotFound)
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}
	log.Println("LOG2")
	cookie, err := info.cookieApp.GenerateCookie()
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI),
			zap.Int("for user", user.UserID),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("LOG3")
	err = info.cookieApp.AddCookieInfo(&entity.CookieInfo{user.UserID, cookie})
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI),
			zap.Int("for user", user.UserID),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("LOG4")
	http.SetCookie(w, cookie)

	// Replacing token in websocket connection info
	token := r.Header.Get("X-CSRF-Token")
	err = info.websocketApp.ChangeToken(user.UserID, token)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI),
			zap.Int("for user", user.UserID),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HandleLogoutUser logs current user out of their session
func (info *AuthInfo) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	userCookie := r.Context().Value(entity.CookieInfoKey).(*entity.CookieInfo)

	err := info.cookieApp.RemoveCookie(userCookie)
	if err != nil {
		info.logger.Info(err.Error(), zap.String("url", r.RequestURI),
			zap.Int("for user", userCookie.UserID),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCookie.Cookie.Expires = time.Now().AddDate(0, 0, -1) // Making cookie expire
	http.SetCookie(w, userCookie.Cookie)

	w.WriteHeader(http.StatusNoContent)
}

// HandleCheckUser checks if current user is logged in
func (info *AuthInfo) HandleCheckUser(w http.ResponseWriter, r *http.Request) {
	_, found := middleware.CheckCookies(r, info.authApp)
	if !found {
		info.logger.Info(entity.UnauthorizedError.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
