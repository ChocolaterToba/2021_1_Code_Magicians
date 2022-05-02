package repository

import (
	"context"
	"pinterest/services/product/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepoInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	UpdateShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	UpdateProduct(ctx context.Context, product domain.Product) (err error)
	GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
}

type ProductRepo struct {
	postgresDB *pgxpool.Pool
}

func NewProductRepo(postgresDB *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{postgresDB: postgresDB}
}
