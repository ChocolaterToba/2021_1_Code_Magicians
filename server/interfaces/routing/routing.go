package routing

import (
	"net/http"
	"os"
	authclient "pinterest/clients/auth"
	authfacade "pinterest/interfaces/auth"
	"pinterest/interfaces/metrics"
	mid "pinterest/interfaces/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func CreateRouter(authClient authclient.AuthClientInterface, authFacade *authfacade.AuthFacade, csrfOn bool, httpOn bool) *mux.Router {
	r := mux.NewRouter()

	r.Use(mid.PanicMid, metrics.PrometheusMiddleware)

	if csrfOn {
		csrfMid := csrf.Protect(
			[]byte(os.Getenv("CSRF_KEY")),
			csrf.Path("/"),
			csrf.Secure(httpOn), // REMOVE IN PROD!!!!
		)
		r.Use(csrfMid)
		r.Use(mid.CSRFSettingMid)
	}

	r.Handle("/metrics", promhttp.Handler())

	// r.HandleFunc("/api/auth/signup", mid.NoAuthMid(authInfo.HandleCreateUser, authApp)).Methods("POST")
	r.HandleFunc("/api/auth/login", mid.NoAuthMid(authFacade.LoginUser, authClient)).Methods("POST")
	r.HandleFunc("/api/auth/logout", mid.AuthMid(authFacade.LogoutUser, authClient)).Methods("POST")
	r.HandleFunc("/api/auth/check", authFacade.CheckUser).Methods("GET")

	// r.HandleFunc("/api/profile/password", mid.AuthMid(profileInfo.HandleChangePassword, authApp)).Methods("PUT")
	// r.HandleFunc("/api/profile/edit", mid.AuthMid(profileInfo.HandleEditProfile, authApp)).Methods("PUT")
	// r.HandleFunc("/api/profile/delete", mid.AuthMid(profileInfo.HandleDeleteProfile, authApp)).Methods("DELETE")
	// r.HandleFunc("/api/profile/{id:[0-9]+}", profileInfo.HandleGetProfile).Methods("GET") // Is preferred over next one
	// r.HandleFunc("/api/profile/{username}", profileInfo.HandleGetProfile).Methods("GET")
	// r.HandleFunc("/api/profile", mid.AuthMid(profileInfo.HandleGetProfile, authApp)).Methods("GET")
	// r.HandleFunc("/api/profile/avatar", mid.AuthMid(profileInfo.HandlePostAvatar, authApp)).Methods("PUT")
	// r.HandleFunc("/api/profiles/search/{searchKey}", profileInfo.HandleGetProfilesByKeyWords).Methods("GET")

	if csrfOn {
		r.HandleFunc("/api/csrf", func(w http.ResponseWriter, r *http.Request) { // Is used only for getting csrf key
			w.WriteHeader(http.StatusCreated)
		}).Methods("GET")
	}

	return r
}
