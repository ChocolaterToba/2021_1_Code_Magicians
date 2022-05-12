package product

import (
	"context"
	"encoding/json"
	"net/http"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Get order ...
func (facade *ProductFacade) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	vars := mux.Vars(r)
	orderIDStr, passedID := vars[string(domain.IDKey)]

	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderID, _ := strconv.ParseUint(orderIDStr, 10, 64)
	order, err := facade.productClient.GetOrderByID(context.Background(), cookie.UserID, orderID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		switch err {
		case domain.ErrOrderNotFound:
			w.WriteHeader(http.StatusNotFound)
			return
		case domain.ErrForeignOrder:
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(order)
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

func (facade *ProductFacade) GetOrders(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value(domain.CookieInfoKey).(domain.CookieInfo)

	orders, err := facade.productClient.GetOrders(context.Background(), cookie.UserID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(orders)
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
