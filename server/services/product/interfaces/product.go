package facade

import (
	"context"
	"pinterest/services/product/application"
	"pinterest/services/product/domain"
	pb "pinterest/services/product/proto"

	"github.com/pkg/errors"
	_ "google.golang.org/grpc"
)

type ProductFacade struct {
	pb.UnimplementedProductServiceServer
	app application.ProductAppInterface
}

func NewProductFacade(app application.ProductAppInterface) *ProductFacade {
	return &ProductFacade{
		app: app,
	}
}

func (facade *ProductFacade) CreateShop(ctx context.Context, in *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
	shopInput := domain.Shop{
		Title:       in.Title,
		Description: in.Description,
		ManagerIDs:  in.ManagerIds,
	}

	shopID, err := facade.app.CreateShop(ctx, shopInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create shop:")
	}

	return &pb.CreateShopResponse{Id: shopID}, nil
}

func (facade *ProductFacade) EditShop(ctx context.Context, in *pb.EditShopRequest) (*pb.Empty, error) {
	shopInput := domain.Shop{
		Id:          in.Id,
		Title:       in.Title,
		Description: in.Description,
		ManagerIDs:  in.ManagerIds,
	}

	err := facade.app.EditShop(ctx, shopInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not edit shop:")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) GetShopByID(ctx context.Context, in *pb.GetShopRequest) (*pb.Shop, error) {
	shop, err := facade.app.GetShopByID(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get shop:")
	}

	return domain.ToPbShop(shop), nil
}

func (facade *ProductFacade) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productInput := domain.Product{
		Id:           0,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Size:         in.Size,
		Category:     in.Category,
		ShopId:       in.ShopId,
	}

	productID, err := facade.app.CreateProduct(ctx, productInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create product:")
	}

	return &pb.CreateProductResponse{Id: productID}, nil
}

func (facade *ProductFacade) EditProduct(ctx context.Context, in *pb.EditProductRequest) (*pb.Empty, error) {
	productInput := domain.Product{
		Id:           in.Id,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Size:         in.Size,
		Category:     in.Category,
		ShopId:       in.ShopId,
	}

	err := facade.app.EditProduct(ctx, productInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not edit product:")
	}

	return &pb.Empty{}, nil
}
func (facade *ProductFacade) GetProductByID(ctx context.Context, in *pb.GetProductRequest) (*pb.Product, error) {
	product, err := facade.app.GetProductByID(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get product:")
	}

	return domain.ToPbProduct(product), nil
}

func (facade *ProductFacade) GetProductsByIDs(ctx context.Context, in *pb.GetProductsByIDsRequest) (*pb.GetProductsResponse, error) {
	products, err := facade.app.GetProductsByIDs(ctx, in.GetIds())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products:")
	}

	return &pb.GetProductsResponse{
		Products: domain.ToPbProducts(products),
	}, nil
}

func (facade *ProductFacade) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	products, err := facade.app.GetProducts(ctx, in.GetPageOffset(), in.GetPageSize(), in.GetCategory())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get products:")
	}

	return &pb.GetProductsResponse{
		Products: domain.ToPbProducts(products),
	}, nil
}

func (facade *ProductFacade) AddToCart(ctx context.Context, in *pb.AddToCartRequest) (*pb.Empty, error) {
	err := facade.app.AddToCart(ctx, in.GetUserId(), in.GetProductId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not add product to cart:")
	}

	return &pb.Empty{}, nil
}

func (facade *ProductFacade) GetCart(ctx context.Context, in *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cart, err := facade.app.GetCart(ctx, in.GetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get cart:")
	}

	productIds := make([]uint64, 0) // is needed because batch-getting products is faster than one-by-one solution
	for id := range cart {
		productIds = append(productIds, id)
	}

	products, err := facade.app.GetProductsByIDs(ctx, productIds)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get cart:")
	}

	result := make([]*pb.ProductWithQuantity, len(products))
	for _, product := range products {
		result = append(result, &pb.ProductWithQuantity{
			Product:  domain.ToPbProduct(product),
			Quantity: cart[product.Id],
		})
	}

	return &pb.GetCartResponse{
		Products: result,
	}, nil
}

func (facade *ProductFacade) CompleteCart(ctx context.Context, in *pb.CompleteCartRequest) (*pb.Empty, error) {
	err := facade.app.CompleteCart(ctx, in.GetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not complete cart:")
	}

	return &pb.Empty{}, nil
}
