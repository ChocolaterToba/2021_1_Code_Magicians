package profile

import (
	"context"
	"net/http"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// AddTocart adds product with assed productID to current user's cart
func (facade *ProductFacade) AddToCart(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	vars := mux.Vars(r)
	productIDStr, passedID := vars[string(domain.IDKey)]

	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productID, _ := strconv.ParseUint(productIDStr, 10, 64)

	err := facade.productClient.AddToCart(context.Background(), cookie.UserID, productID)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrProductNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
