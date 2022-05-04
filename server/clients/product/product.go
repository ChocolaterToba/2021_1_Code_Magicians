package product

import (
	"context"
	"pinterest/domain"
	productdomain "pinterest/services/product/domain"
	productproto "pinterest/services/product/proto"
	"strings"

	"github.com/pkg/errors"
)

type ProductClientInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	EditShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	EditProduct(ctx context.Context, product domain.Product) (err error)
	GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
	GetProducts(ctx context.Context, offset uint64, pageSize uint64) (products []domain.Product, err error)
}

type ProductClient struct {
	productClient productproto.ProductServiceClient
}

func NewProductClient(productClient productproto.ProductServiceClient) *ProductClient {
	return &ProductClient{
		productClient: productClient,
	}
}

func (client *ProductClient) CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error) {
	pbShopID, err := client.productClient.CreateShop(ctx, domain.ToPbCreateShopRequest(shop))

	if err != nil {
		return 0, errors.Wrap(err, "product client error: ")
	}

	return pbShopID.Id, nil
}

func (client *ProductClient) EditShop(ctx context.Context, shop domain.Shop) (err error) {
	_, err = client.productClient.EditShop(ctx, domain.ToPbEditShopRequest(shop))

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ShopNotFoundError.Error()) {
			return domain.ErrShopNotFound
		}
		return errors.Wrap(err, "product client error: ")
	}

	return nil
}

func (client *ProductClient) GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error) {
	pbShop, err := client.productClient.GetShopByID(ctx, &productproto.GetShopRequest{Id: id})

	if err != nil {
		return domain.Shop{}, errors.Wrap(err, "product client error: ")
	}

	return domain.ToShop(pbShop), nil
}

func (client *ProductClient) CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error) {
	pbProductID, err := client.productClient.CreateProduct(ctx, domain.ToPbCreateProductRequest(product))

	if err != nil {
		return 0, errors.Wrap(err, "product client error: ")
	}

	return pbProductID.Id, nil
}

func (client *ProductClient) EditProduct(ctx context.Context, product domain.Product) (err error) {
	_, err = client.productClient.EditProduct(ctx, domain.ToPbEditProductRequest(product))

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ProductNotFoundError.Error()) {
			return domain.ErrProductNotFound
		}
		return errors.Wrap(err, "product client error: ")
	}

	return nil
}

func (client *ProductClient) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	pbProduct, err := client.productClient.GetProductByID(ctx, &productproto.GetProductRequest{Id: id})

	if err != nil {
		return domain.Product{}, errors.Wrap(err, "product client error: ")
	}

	return domain.ToProduct(pbProduct), nil
}

func (client *ProductClient) GetProducts(ctx context.Context, offset uint64, pageSize uint64) (products []domain.Product, err error) {
	pbProducts, err := client.productClient.GetProducts(ctx,
		&productproto.GetProductsRequest{
			Offset:   offset,
			PageSize: pageSize,
		})

	if err != nil {
		return nil, errors.Wrap(err, "product client error: ")
	}

	return domain.ToProducts(pbProducts.Products), nil
}
