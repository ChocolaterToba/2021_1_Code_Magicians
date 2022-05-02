package application

import (
	"context"
	"pinterest/services/product/domain"
	repository "pinterest/services/product/infrastructure"

	"github.com/pkg/errors"
)

type ProductAppInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	EditShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	EditProduct(ctx context.Context, product domain.Product) (err error)
	GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
}

type ProductApp struct {
	repo repository.ProductRepoInterface
}

func NewProductApp(repo repository.ProductRepoInterface) *ProductApp {
	return &ProductApp{
		repo: repo,
	}
}

func (app *ProductApp) CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error) {
	// TODO: check if current user's ID is in shop.ManagerIDs
	return app.repo.CreateShop(ctx, shop)
}

func (app *ProductApp) EditShop(ctx context.Context, shop domain.Shop) (err error) {
	//TODO: add transactions here?
	dbShop, err := app.repo.GetShopByID(ctx, shop.Id)
	if err != nil {
		return err
	}

	if shop.Title != "" {
		dbShop.Title = shop.Title
	}
	if shop.Description != "" {
		dbShop.Description = shop.Description
	}
	if len(shop.ManagerIDs) != 0 {
		// TODO: check if current user's id is in managerIDs
		dbShop.ManagerIDs = shop.ManagerIDs
	}

	return app.repo.UpdateShop(ctx, dbShop)
}

func (app *ProductApp) GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error) {
	return app.repo.GetShopByID(ctx, id)
}

func (app *ProductApp) CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error) {
	return 0, errors.New("Not implemented yet")
}

func (app *ProductApp) EditProduct(ctx context.Context, product domain.Product) (err error) {
	return errors.New("Not implemented yet")
}

func (app *ProductApp) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	return domain.Product{}, errors.New("Not implemented yet")
}
