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
	// CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	// EditProduct(ctx context.Context, product domain.Product) (err error)
	// GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
}

type ProductClient struct {
	productClient productproto.ProductClient
}

func NewProductClient(productClient productproto.ProductClient) *ProductClient {
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

	return *domain.ToShop(pbShop), nil
}
