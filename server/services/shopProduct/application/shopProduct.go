package application

import (
	"context"
	"pinterest/services/shopProduct/domain"
)

type ShopProductAppInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	EditShop(ctx context.Context, shop domain.Shop) (err error)
	GetShop(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	EditProduct(ctx context.Context, product domain.Product) (err error)
	GetProduct(ctx context.Context, id uint64) (product domain.Product, err error)
}
