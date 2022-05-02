package application

import (
	"context"
	"pinterest/services/product/domain"
	repository "pinterest/services/product/infrastructure"
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

	dbShop.Title = replaceStringIfNotEmpty(dbShop.Title, shop.Title)
	dbShop.Description = replaceStringIfNotEmpty(dbShop.Description, shop.Description)
	dbShop.ManagerIDs = replaceSliceIfNotEmpty(dbShop.ManagerIDs, shop.ManagerIDs)

	return app.repo.UpdateShop(ctx, dbShop)
}

func (app *ProductApp) GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error) {
	return app.repo.GetShopByID(ctx, id)
}

func (app *ProductApp) CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error) {
	// TODO: check if current user's ID is in shop.ManagerIDs
	return app.repo.CreateProduct(ctx, product)
}

func (app *ProductApp) EditProduct(ctx context.Context, product domain.Product) (err error) {
	//TODO: add transactions here?
	dbProduct, err := app.repo.GetProductByID(ctx, product.Id)
	if err != nil {
		return err
	}

	dbProduct.Title = replaceStringIfNotEmpty(dbProduct.Title, product.Title)
	dbProduct.Description = replaceStringIfNotEmpty(dbProduct.Description, product.Description)
	dbProduct.Price = replaceUint64IfNotEmpty(dbProduct.Price, product.Price)
	dbProduct.Availability = product.Availability // there isn't really "empty" bool value
	dbProduct.AssemblyTime = replaceUint64IfNotEmpty(dbProduct.AssemblyTime, product.AssemblyTime)
	dbProduct.PartsAmount = replaceUint64IfNotEmpty(dbProduct.PartsAmount, product.PartsAmount)
	dbProduct.Size = replaceStringIfNotEmpty(dbProduct.Size, product.Size)
	dbProduct.Category = replaceStringIfNotEmpty(dbProduct.Category, product.Category)
	dbProduct.ShopId = replaceUint64IfNotEmpty(dbProduct.ShopId, product.ShopId)

	return app.repo.UpdateProduct(ctx, dbProduct)
}

func (app *ProductApp) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	return app.repo.GetProductByID(ctx, product.Id)
}

func replaceStringIfNotEmpty(original string, replacement string) (result string) {
	if replacement != "" {
		return replacement
	}

	return original
}

func replaceUint64IfNotEmpty(original uint64, replacement uint64) (result uint64) {
	if replacement != 0 {
		return replacement
	}

	return original
}

func replaceSliceIfNotEmpty(original []uint64, replacement []uint64) (result []uint64) {
	if len(replacement) != 0 {
		return replacement
	}

	return original
}
