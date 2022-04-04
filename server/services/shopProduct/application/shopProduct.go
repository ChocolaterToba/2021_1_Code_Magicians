package application

import (
	"context"
	"pinterest/services/shopProduct/domain"
)

type ShopProductAppInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
}
