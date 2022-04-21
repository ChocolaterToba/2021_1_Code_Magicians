package facade

import (
	"context"
	"pinterest/services/shopProduct/application"
	"pinterest/services/shopProduct/domain"
	pb "pinterest/services/shopProduct/proto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	_ "google.golang.org/grpc"
)

type ShopProductFacade struct {
	postgresDB *pgxpool.Pool
	app        application.ShopProductAppInterface
}

func NewService(postgresDB *pgxpool.Pool, app application.ShopProductAppInterface) *ShopProductFacade {
	return &ShopProductFacade{
		postgresDB: postgresDB,
		app:        app,
	}
}

func (facade *ShopProductFacade) CreateShop(ctx context.Context, in *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
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

func (facade *ShopProductFacade) EditShop(ctx context.Context, in *pb.EditShopRequest) (*pb.Empty, error) {
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

func (facade *ShopProductFacade) GetShop(ctx context.Context, in *pb.GetShopRequest) (*pb.Shop, error) {
	shop, err := facade.app.GetShop(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get shop:")
	}

	return domain.ToPbShop(shop), nil
}

func (facade *ShopProductFacade) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productInput := domain.Product{
		Id:           0,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Rating:       in.Rating,
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

func (facade *ShopProductFacade) EditProduct(ctx context.Context, in *pb.EditProductRequest) (*pb.Empty, error) {
	productInput := domain.Product{
		Id:           in.Id,
		Title:        in.Title,
		Description:  in.Description,
		Price:        in.Price,
		Availability: in.Availability,
		AssemblyTime: in.AssemblyTime,
		PartsAmount:  in.PartsAmount,
		Rating:       in.Rating,
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
func (facade *ShopProductFacade) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*pb.Product, error) {
	product, err := facade.app.GetProduct(ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "Could not get product:")
	}

	return domain.ToPbProduct(product), nil
}
