package profile

import (
	productclient "pinterest/clients/product"

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
