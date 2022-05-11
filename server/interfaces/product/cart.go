package profile

import (
	"context"
	"encoding/json"
	"net/http"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// AddTocart adds product with passed productID to current user's cart
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
	return
}

// RemoveFromCart removes product with passed productID from current user's cart
func (facade *ProductFacade) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
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

	err := facade.productClient.RemoveFromCart(context.Background(), cookie.UserID, productID)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrProductNotFound, domain.ErrProductNotFoundInCart:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

// GetCart returns current user's cart as slice of products and their order amounts
func (facade *ProductFacade) GetCart(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	cart, err := facade.productClient.GetCart(context.Background(), cookie.UserID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cart == nil {
		cart = make([]domain.ProductWithQuantity, 0) // so that cart marshalls as  "[]", not "null"
	}

	responseBody, err := json.Marshal(cart)
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

// CompleteCart sends user's cart to manager and clears it
func (facade *ProductFacade) CompleteCart(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	orderInput := new(domain.Order)
	err := json.NewDecoder(r.Body).Decode(orderInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = facade.productClient.CompleteCart(context.Background(), cookie.UserID, orderInput.PickUp, orderInput.DeliveryAddress, orderInput.PaymentMethod, orderInput.CallNeeded)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrCartEmpty:
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
