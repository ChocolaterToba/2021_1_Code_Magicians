package facade

import (
	"context"
	"pinterest/services/shopProduct/application"
	"pinterest/services/shopProduct/domain"
	pb "pinterest/services/shopProduct/proto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

func (facade *ShopProductFacade) CreateShop(ctx context.Context, in *pb.CreateShopRequest, opts ...grpc.CallOption) (*pb.StatusResponse, error) {
	shopInput := domain.Shop{
		Title:       in.Title,
		Description: in.Description,
		ManagerIDs:  in.ManagerIds,
	}

	_, err := facade.app.CreateShop(ctx, shopInput)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create shop:")
	}

	return &pb.StatusResponse{
		Code:   200,
		Status: "success",
	}, nil
}
