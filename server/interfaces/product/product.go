package profile

import (
	"context"
	"encoding/json"
	"net/http"
	productclient "pinterest/clients/product"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// ProductFacade calls product app
type ProductFacade struct {
	productClient productclient.ProductClientInterface
	logger        *zap.Logger
}

func NewProductFacade(productClient productclient.ProductClientInterface, logger *zap.Logger) *ProductFacade {
	return &ProductFacade{
		productClient: productClient,
		logger:        logger,
	}
}

// Create shop creates shop using provided data
func (facade *ProductFacade) CreateShop(w http.ResponseWriter, r *http.Request) {
	shopInput := new(domain.Shop)
	err := json.NewDecoder(r.Body).Decode(shopInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shopID, err := facade.productClient.CreateShop(context.Background(), *shopInput)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userOutput := domain.ShopIDResponse{ShopID: shopID}
	responseBody, err := json.Marshal(userOutput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// Edit shop ...
func (facade *ProductFacade) EditShop(w http.ResponseWriter, r *http.Request) {
	shopInput := new(domain.Shop)
	err := json.NewDecoder(r.Body).Decode(shopInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	shopIDstr, passedID := vars[string(domain.IDKey)]
	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shopInput.Id, _ = strconv.ParseUint(shopIDstr, 10, 64)

	err = facade.productClient.EditShop(context.Background(), *shopInput)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrShopNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
