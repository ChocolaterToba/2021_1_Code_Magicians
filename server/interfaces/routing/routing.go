package routing

import (
	"net/http"
	"os"
	authclient "pinterest/clients/auth"
	authfacade "pinterest/interfaces/auth"
	"pinterest/interfaces/metrics"
	mid "pinterest/interfaces/middleware"
	productfacade "pinterest/interfaces/product"
	profilefacade "pinterest/interfaces/profile"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func CreateRouter(authClient authclient.AuthClientInterface,
	authFacade *authfacade.AuthFacade, profileFacade *profilefacade.ProfileFacade, productFacade *productfacade.ProductFacade,
	csrfOn bool) *mux.Router {
	r := mux.NewRouter()

	r.Use(mid.PanicMid, metrics.PrometheusMiddleware)

	if csrfOn {
		csrfMid := csrf.Protect(
			[]byte(os.Getenv("CSRF_KEY")),
			csrf.Path("/"),
			csrf.Secure(true),
		)
		r.Use(csrfMid)
		r.Use(mid.CSRFSettingMid)
	}

	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/api/auth/signup", mid.NoAuthMid(profileFacade.CreateUser, authClient)).Methods("POST")
	r.HandleFunc("/api/auth/login", mid.NoAuthMid(authFacade.LoginUser, authClient)).Methods("POST")
	r.HandleFunc("/api/auth/logout", mid.AuthMid(authFacade.LogoutUser, authClient)).Methods("POST")
	r.HandleFunc("/api/auth/check", authFacade.CheckUser).Methods("GET")

	r.HandleFunc("/api/auth/credentials/edit", mid.AuthMid(authFacade.ChangeCredentials, authClient)).Methods("PUT")
	r.HandleFunc("/api/profile/edit", mid.AuthMid(profileFacade.EditUser, authClient)).Methods("PUT")
	// r.HandleFunc("/api/profile/delete", mid.AuthMid(profileInfo.HandleDeleteProfile, authApp)).Methods("DELETE")
	r.HandleFunc("/api/profile", mid.AuthMid(profileFacade.GetCurrentUser, authClient)).Methods("GET")
	r.HandleFunc("/api/profile/{id:[0-9]+}", profileFacade.GetUserByID).Methods("GET") // Is preferred over next one
	r.HandleFunc("/api/profile/{username}", profileFacade.GetUserByUsername).Methods("GET")
	// r.HandleFunc("/api/profile/avatar", mid.AuthMid(profileInfo.HandlePostAvatar, authApp)).Methods("PUT")
	// r.HandleFunc("/api/profiles/search/{searchKey}", profileInfo.HandleGetProfilesByKeyWords).Methods("GET")

	r.HandleFunc("/api/shop", mid.AuthMid(productFacade.CreateShop, authClient)).Methods("POST")

	if csrfOn {
		r.HandleFunc("/api/csrf", func(w http.ResponseWriter, r *http.Request) { // Is used only for getting csrf key
			w.WriteHeader(http.StatusCreated)
		}).Methods("GET")
	}

	return r
}
