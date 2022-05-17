package product

import (
	"bufio"
	"context"
	"io"
	"pinterest/domain"
	productdomain "pinterest/services/product/domain"
	productproto "pinterest/services/product/proto"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type ProductClientInterface interface {
	CreateShop(ctx context.Context, shop domain.Shop) (id uint64, err error)
	EditShop(ctx context.Context, shop domain.Shop) (err error)
	GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error)
	CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error)
	EditProduct(ctx context.Context, product domain.Product) (err error)
	UpdateProductAvatars(ctx context.Context, productID uint64, avatars []domain.FileWithName) (err error)
	UpdateProductVideo(ctx context.Context, productID uint64, filename string, file io.Reader) (err error)
	GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error)
	GetProductsByIDs(ctx context.Context, ids []uint64) (products []domain.Product, err error)
	GetProducts(ctx context.Context, pageOffset uint64, pageSize uint64, category string) (products []domain.Product, err error)
	AddToCart(ctx context.Context, userID uint64, productID uint64) (err error)
	RemoveFromCart(ctx context.Context, userID uint64, productID uint64) (err error)
	GetCart(ctx context.Context, userID uint64) (cart []domain.ProductWithQuantity, err error)
	CompleteCart(ctx context.Context, userID uint64, pickUp bool, deliveryAddress string, paymentMethod string, callNeeded bool) (err error)
	GetOrderByID(ctx context.Context, userID uint64, orderID uint64) (order domain.Order, err error)
	GetOrders(ctx context.Context, userID uint64) (orders []domain.Order, err error)
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
		return 0, errors.Wrap(err, "product client error")
	}

	return pbShopID.Id, nil
}

func (client *ProductClient) EditShop(ctx context.Context, shop domain.Shop) (err error) {
	_, err = client.productClient.EditShop(ctx, domain.ToPbEditShopRequest(shop))

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ShopNotFoundError.Error()) {
			return domain.ErrShopNotFound
		}
		return errors.Wrap(err, "product client error")
	}

	return nil
}

func (client *ProductClient) GetShopByID(ctx context.Context, id uint64) (shop domain.Shop, err error) {
	pbShop, err := client.productClient.GetShopByID(ctx, &productproto.GetShopRequest{Id: id})

	if err != nil {
		return domain.Shop{}, errors.Wrap(err, "product client error")
	}

	return domain.ToShop(pbShop), nil
}

func (client *ProductClient) CreateProduct(ctx context.Context, product domain.Product) (id uint64, err error) {
	pbProductID, err := client.productClient.CreateProduct(ctx, domain.ToPbCreateProductRequest(product))

	if err != nil {
		return 0, errors.Wrap(err, "product client error")
	}

	return pbProductID.Id, nil
}

func (client *ProductClient) EditProduct(ctx context.Context, product domain.Product) (err error) {
	_, err = client.productClient.EditProduct(ctx, domain.ToPbEditProductRequest(product))

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ProductNotFoundError.Error()) {
			return domain.ErrProductNotFound
		}
		return errors.Wrap(err, "product client error")
	}

	return nil
}

func (client *ProductClient) UpdateProductAvatars(ctx context.Context, productID uint64, avatars []domain.FileWithName) (err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := client.productClient.UpdateProductAvatars(ctx)
	if err != nil {
		return errors.Wrap(err, "Cannot start stream")
	}

	productIDReq := &productproto.UpdateProductAvatarsRequest{
		Data: &productproto.UpdateProductAvatarsRequest_ProductId{
			ProductId: productID,
		},
	}
	err = stream.Send(productIDReq)
	if err != nil {
		return errors.Wrap(err, "Cannot send product id to service")
	}

	for _, avatar := range avatars {
		filenameReq := &productproto.UpdateProductAvatarsRequest{
			Data: &productproto.UpdateProductAvatarsRequest_Filename{
				Filename: avatar.Filename,
			},
		}
		err = stream.Send(filenameReq)
		if err != nil {
			return errors.Wrap(err, "Cannot send image's filename to service")
		}

		reader := bufio.NewReader(avatar.File)
		buffer := make([]byte, 3.5*1024*1024) // jrpc by default cannot receive packages larger than 4 MB

		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return errors.Wrap(err, "Cannot read chunk to buffer")
			}

			req := &productproto.UpdateProductAvatarsRequest{
				Data: &productproto.UpdateProductAvatarsRequest_Chunk{
					Chunk: buffer[:n],
				},
			}
			err = stream.Send(req)
			if err != nil {
				return errors.Wrap(err, "Cannot send chunk to server")
			}
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		// TODO: parse this error
		return errors.Wrap(err, "Cannot receive response")
	}

	return nil
}

func (client *ProductClient) UpdateProductVideo(ctx context.Context, productID uint64, filename string, file io.Reader) (err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := client.productClient.UpdateProductVideo(ctx)
	if err != nil {
		return errors.Wrap(err, "Cannot start stream")
	}

	productIDReq := &productproto.UpdateProductVideoRequest{
		Data: &productproto.UpdateProductVideoRequest_ProductId{
			ProductId: productID,
		},
	}
	err = stream.Send(productIDReq)
	if err != nil {
		return errors.Wrap(err, "Cannot send product id to service")
	}

	filenameReq := &productproto.UpdateProductVideoRequest{
		Data: &productproto.UpdateProductVideoRequest_Filename{
			Filename: filename,
		},
	}
	err = stream.Send(filenameReq)
	if err != nil {
		return errors.Wrap(err, "Cannot send video's filename to service")
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 3.5*1024*1024) // jrpc by default cannot receive packages larger than 4 MB

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "Cannot read chunk to buffer")
		}

		req := &productproto.UpdateProductVideoRequest{
			Data: &productproto.UpdateProductVideoRequest_Chunk{
				Chunk: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			return errors.Wrap(err, "Cannot send chunk to server")
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		// TODO: parse this error
		return errors.Wrap(err, "Cannot receive response")
	}

	return nil
}

func (client *ProductClient) GetProductByID(ctx context.Context, id uint64) (product domain.Product, err error) {
	pbProduct, err := client.productClient.GetProductByID(ctx, &productproto.GetProductRequest{Id: id})

	if err != nil {
		return domain.Product{}, errors.Wrap(err, "product client error")
	}

	return domain.ToProduct(pbProduct), nil
}

func (client *ProductClient) GetProductsByIDs(ctx context.Context, ids []uint64) (products []domain.Product, err error) {
	pbProducts, err := client.productClient.GetProductsByIDs(ctx, &productproto.GetProductsByIDsRequest{Ids: ids})
	if err != nil {
		if strings.Contains(err.Error(), productdomain.ProductNotFoundError.Error()) {
			return nil, errors.Wrap(err, domain.ErrProductNotFound.Error()) // We return "raw" err here so that it can be given to client later
		}
		return nil, errors.Wrap(err, "product client error")
	}

	return domain.ToProducts(pbProducts.Products), nil
}

func (client *ProductClient) GetProducts(ctx context.Context, pageOffset uint64, pageSize uint64, category string) (products []domain.Product, err error) {
	pbProducts, err := client.productClient.GetProducts(ctx,
		&productproto.GetProductsRequest{
			PageOffset: pageOffset,
			PageSize:   pageSize,
			Category:   category,
		})

	if err != nil {
		return nil, errors.Wrap(err, "product client error")
	}

	return domain.ToProducts(pbProducts.Products), nil
}

func (client *ProductClient) AddToCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	_, err = client.productClient.AddToCart(ctx, &productproto.AddToCartRequest{
		UserId:    userID,
		ProductId: productID,
	})

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ProductNotFoundError.Error()) {
			return domain.ErrProductNotFound
		}
		return errors.Wrap(err, "product client error")
	}

	return nil
}

func (client *ProductClient) RemoveFromCart(ctx context.Context, userID uint64, productID uint64) (err error) {
	_, err = client.productClient.RemoveFromCart(ctx, &productproto.RemoveFromCartRequest{
		UserId:    userID,
		ProductId: productID,
	})

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ProductNotFoundInCartError.Error()) {
			return domain.ErrProductNotFoundInCart
		}
		if strings.Contains(err.Error(), productdomain.ProductNotFoundError.Error()) {
			return domain.ErrProductNotFound
		}
		return errors.Wrap(err, "product client error")
	}

	return nil
}

func (client *ProductClient) GetCart(ctx context.Context, userID uint64) (cart []domain.ProductWithQuantity, err error) {
	pbCart, err := client.productClient.GetCart(ctx, &productproto.GetCartRequest{
		UserId: userID,
	})

	if err != nil {
		return nil, errors.Wrap(err, "product client error")
	}

	return domain.ToProductsWithQuantity(pbCart.Products), nil
}

func (client *ProductClient) CompleteCart(ctx context.Context, userID uint64, pickUp bool, deliveryAddress string, paymentMethod string, callNeeded bool) (err error) {
	_, err = client.productClient.CompleteCart(ctx, &productproto.CompleteCartRequest{
		UserId:          userID,
		PickUp:          pickUp,
		DeliveryAddress: deliveryAddress,
		PaymentMethod:   paymentMethod,
		CallNeeded:      callNeeded,
	})

	if err != nil {
		if strings.Contains(err.Error(), productdomain.CartEmptyError.Error()) {
			return domain.ErrCartEmpty
		}
		return errors.Wrap(err, "product client error")
	}

	return nil
}

func (client *ProductClient) GetOrderByID(ctx context.Context, userID uint64, orderID uint64) (order domain.Order, err error) {
	pbOrder, err := client.productClient.GetOrderByID(ctx, &productproto.GetOrderByIDRequest{
		OrderId: orderID,
		UserId:  userID,
	})

	if err != nil {
		if strings.Contains(err.Error(), productdomain.ForeignOrderError.Error()) {
			return domain.Order{}, domain.ErrForeignOrder
		}
		if strings.Contains(err.Error(), productdomain.OrderNotFoundError.Error()) {
			return domain.Order{}, domain.ErrOrderNotFound
		}

		return domain.Order{}, err
	}

	return domain.ToOrder(pbOrder), nil
}

func (client *ProductClient) GetOrders(ctx context.Context, userID uint64) (orders []domain.Order, err error) {
	pbOrders, err := client.productClient.GetOrders(ctx, &productproto.GetOrdersRequest{
		UserId: userID,
	})

	if err != nil {
		return nil, err
	}

	return domain.ToOrders(pbOrders.Orders), nil
}
