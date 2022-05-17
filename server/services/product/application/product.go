package application

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	s3client "pinterest/services/product/clients/s3"
	"pinterest/services/product/domain"
	repository "pinterest/services/product/infrastructure"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/pkg/errors"
)

type ProductAppInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	EditShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	EditProduct(ctx context.Context, product domain.Product) (err error)
	UpdateProductAvatars(ctx context.Context, productID uint64, avatars []domain.FileWithName) (err error)
	UpdateProductVideo(ctx context.Context, productID uint64, filename string, file *bytes.Buffer) (err error)
	GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
	GetProductsByIDs(ctx context.Context, ids []uint64) (products []domain.Product, err error)
	GetProducts(ctx context.Context, pageOffset uint64, pageSize uint64, category string) (products []domain.Product, err error)
	AddToCart(ctx context.Context, userID uint64, productID uint64) (err error)
	RemoveFromCart(ctx context.Context, userID uint64, productID uint64) (err error)
	GetCart(ctx context.Context, userID uint64) (cart map[uint64]uint64, err error)
	CompleteCart(ctx context.Context, orderBillet domain.Order) (err error)
	GetOrderByID(ctx context.Context, userID uint64, orderID uint64) (order domain.Order, err error)
	GetOrdersByUserID(ctx context.Context, userID uint64) (orders []domain.Order, err error)
}

type ProductApp struct {
	repo     repository.ProductRepoInterface
	s3Client s3client.S3ClientInterface
}

func NewProductApp(repo repository.ProductRepoInterface, s3Client s3client.S3ClientInterface) *ProductApp {
	return &ProductApp{
		repo:     repo,
		s3Client: s3Client,
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

const AvatarIDLen = 10

func (app *ProductApp) UpdateProductAvatars(ctx context.Context, productID uint64, avatars []domain.FileWithName) (err error) {
	product, err := app.repo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	oldFilenames := product.ImageLinks
	product.ImageLinks = nil
	for _, avatar := range avatars {
		extension := filepath.Ext(avatar.Filename)
		if extension != ".jpeg" && extension != ".jpg" && extension != ".png" && extension != ".mp4" {
			return errors.Wrap(domain.UnsupportedExtensionError, fmt.Sprintf("Extension: %s", extension))
		}

		filePrefix := time.Now().Format("2006/01/02") // is used for easy manual file search

		fileID := uniuri.NewLen(AvatarIDLen)

		newAvatarPath := filePrefix + "/" + fileID + "_" + avatar.Filename
		product.ImageLinks = append(product.ImageLinks, newAvatarPath)

		err = app.s3Client.UploadFile(ctx, newAvatarPath, avatar.File)
		if err != nil {
			return err
		}
	}

	err = app.repo.UpdateProduct(ctx, product)
	if err != nil {
		for _, filename := range product.ImageLinks {
			app.s3Client.DeleteFile(ctx, filename) // Try to delete freshly uploaded files
		}

		return err
	}

	for _, filename := range oldFilenames {
		if filename == "" {
			continue
		}

		err = app.s3Client.DeleteFile(ctx, filename)
		if err != nil {
			return errors.Wrap(err, "Error when deleting old avatars")
		}
	}

	return nil
}

func (app *ProductApp) UpdateProductVideo(ctx context.Context, productID uint64, filename string, file *bytes.Buffer) (err error) {
	product, err := app.repo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	extension := filepath.Ext(filename)
	if extension != ".mp4" {
		return errors.Wrap(domain.UnsupportedExtensionError, fmt.Sprintf("Extension: %s", extension))
	}

	filePrefix := time.Now().Format("2006/01/02") // is used for easy manual file search

	fileID := uniuri.NewLen(AvatarIDLen)

	newVideoLink := filePrefix + "/" + fileID + "_" + filename

	err = app.s3Client.UploadFile(ctx, newVideoLink, file)
	if err != nil {
		return err
	}

	oldVideoLink := product.VideoLink
	product.VideoLink = newVideoLink

	err = app.repo.UpdateProduct(ctx, product)
	if err != nil {
		app.s3Client.DeleteFile(ctx, oldVideoLink) // Try to delete freshly uploaded file
		return err
	}

	if oldVideoLink == "" {
		return nil
	}

	return app.s3Client.DeleteFile(ctx, oldVideoLink)
}

func (app *ProductApp) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	return app.repo.GetProductByID(ctx, id)
}

func (app *ProductApp) GetProductsByIDs(ctx context.Context, ids []uint64) (products []domain.Product, err error) {
	if len(ids) == 0 {
		return nil, nil
	}

	dbProducts, err := app.repo.GetProductsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	productsMap := make(map[uint64]domain.Product)
	for _, product := range dbProducts {
		productsMap[product.Id] = product
	}

	notFoundProducts := make([]uint64, 0)
	products = make([]domain.Product, 0, len(ids))
	for _, id := range ids {
		product, found := productsMap[id]
		if !found {
			notFoundProducts = append(notFoundProducts, id)
		} else {
			products = append(products, product)
		}
	}

	if len(notFoundProducts) != 0 {
		notFoundProductsStrings := make([]string, 0, len(notFoundProducts))
		for _, id := range notFoundProducts {
			notFoundProductsStrings = append(notFoundProductsStrings, fmt.Sprintf("продукт с id %d не найден", id))
		}
		return nil, errors.Wrap(errors.New(strings.Join(notFoundProductsStrings, ", ")), domain.ProductNotFoundError.Error())
	}

	return products, nil
}

func (app *ProductApp) GetProducts(ctx context.Context, pageOffset uint64, pageSize uint64, category string) (products []domain.Product, err error) {
	if pageSize == 0 {
		pageSize = domain.DefaultPageSize
	}

	return app.repo.GetProducts(ctx, pageOffset, pageSize, category)
}

func (app *ProductApp) AddToCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	cart, err := app.repo.GetCart(ctx, userID)
	if err != nil {
		if err == domain.CartNotFoundError {
			_, err := app.repo.CreateCart(ctx, userID)
			if err != nil {
				return err
			}

			cart = make(map[uint64]uint64) // emptying cart just in case
		} else {
			return err
		}
	}

	// check if product exists
	_, err = app.repo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	cart[productID] = cart[productID] + 1
	return app.repo.UpdateCart(ctx, userID, cart)
}

func (app *ProductApp) RemoveFromCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	cart, err := app.repo.GetCart(ctx, userID)
	if err != nil {
		if err != domain.CartNotFoundError {
			return err
		}

		_, err := app.repo.CreateCart(ctx, userID)
		if err != nil {
			return err
		}

		return domain.ProductNotFoundInCartError // can't remove product from newly created cart
	}

	// check if product exists
	_, err = app.repo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	if cart[productID] < 1 {
		return domain.ProductNotFoundInCartError
	}

	cart[productID] = cart[productID] - 1
	if cart[productID] == 0 {
		delete(cart, productID)
	}
	return app.repo.UpdateCart(ctx, userID, cart)
}

func (app *ProductApp) GetCart(ctx context.Context, userID uint64) (cart map[uint64]uint64, err error) {
	cart, err = app.repo.GetCart(ctx, userID)
	if err != nil {
		if err != domain.CartNotFoundError {
			return nil, err
		}

		_, err := app.repo.CreateCart(ctx, userID)
		if err != nil {
			return nil, err
		}

		cart = make(map[uint64]uint64)
		return cart, nil
	}

	return cart, nil
}

func (app *ProductApp) CompleteCart(ctx context.Context, orderBillet domain.Order) (err error) {
	cart, err := app.repo.GetCart(ctx, orderBillet.UserID)
	if err != nil {
		return err
	}

	if len(cart) == 0 {
		return domain.CartEmptyError
	}

	orderBillet.Items = cart

	productIDs := make([]uint64, len(cart))
	for id := range cart {
		productIDs = append(productIDs, id)
	}

	products, err := app.repo.GetProductsByIDs(ctx, productIDs)
	if err != nil {
		return err
	}

	for _, product := range products {
		orderBillet.TotalPrice = product.Price * cart[product.Id]
	}

	orderBillet.Status = domain.StatusOrderCreated

	fmt.Println("Тут мы якобы шлём письмо, что к нам поступил такой-то заказ, это происходит в отдельной горутине")
	for key := range cart {
		fmt.Printf("id товара %d, количество %d\n", key, cart[key])
	}

	err = app.repo.UpdateCart(ctx, orderBillet.UserID, make(map[uint64]uint64)) // emptying existing cart
	if err != nil {
		return err
	}

	_, err = app.repo.CreateOrder(ctx, orderBillet)
	if err != nil {
		return err
	}

	return nil
}

func (app *ProductApp) GetOrderByID(ctx context.Context, userID uint64, orderID uint64) (order domain.Order, err error) {
	order, err = app.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return domain.Order{}, err
	}

	if order.UserID != userID {
		return domain.Order{}, domain.ForeignOrderError
	}

	return order, nil
}

func (app *ProductApp) GetOrdersByUserID(ctx context.Context, userID uint64) (orders []domain.Order, err error) {
	return app.repo.GetOrdersByUserID(ctx, userID)
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
