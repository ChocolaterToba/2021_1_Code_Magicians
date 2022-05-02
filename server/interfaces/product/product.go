package profile

import (
	"context"
	"encoding/json"
	"net/http"
	productclient "pinterest/clients/product"
	"pinterest/domain"

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
