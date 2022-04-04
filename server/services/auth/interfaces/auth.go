package facade

import (
	"context"
	"pinterest/services/auth/application"
	"pinterest/services/auth/domain"
	pb "pinterest/services/auth/proto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

type AuthFacade struct {
	postgresDB *pgxpool.Pool
	app        application.AuthAppInterface
}

func NewService(postgresDB *pgxpool.Pool, app application.AuthAppInterface) *AuthFacade {
	return &AuthFacade{
		postgresDB: postgresDB,
		app:        app,
	}
}

func (facade *AuthFacade) CheckUserCredentials(ctx context.Context, in *pb.UserAuth, opts ...grpc.CallOption) (*pb.Error, error) {
	_, err := facade.app.CheckUserCredentials(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return &pb.Error{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return &pb.Error{}, nil
}

func (facade *AuthFacade) AddCookieInfo(ctx context.Context, in *pb.CookieInfo, opts ...grpc.CallOption) (*pb.Error, error) {
	err := facade.app.AddCookieInfo(ctx, domain.ToCookieInfo(in))
	if err != nil {
		return &pb.Error{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return &pb.Error{}, nil
}

func (facade *AuthFacade) SearchCookieByValue(ctx context.Context, in *pb.CookieValue, opts ...grpc.CallOption) (*pb.CookieInfo, error) {
	result, err := facade.app.SearchCookieByValue(ctx, in.GetCookieValue())
	if err != nil {
		return &pb.CookieInfo{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return domain.ToPbCookieInfo(result), nil
}

func (facade *AuthFacade) SearchCookieByUserID(ctx context.Context, in *pb.UserID, opts ...grpc.CallOption) (*pb.CookieInfo, error) {
	result, err := facade.app.SearchCookieByUserID(ctx, in.GetUid())
	if err != nil {
		return &pb.CookieInfo{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return domain.ToPbCookieInfo(result), nil
}

func (facade *AuthFacade) RemoveCookie(ctx context.Context, in *pb.CookieInfo, opts ...grpc.CallOption) (*pb.Error, error) {
	err := facade.app.RemoveCookie(ctx, domain.ToCookieInfo(in))
	if err != nil {
		return &pb.Error{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return &pb.Error{}, nil
}

func (facade *AuthFacade) GetUserIDByVkID(ctx context.Context, in *pb.VkIDInfo, opts ...grpc.CallOption) (*pb.UserID, error) {
	userID, err := facade.app.GetUserIDByVkID(ctx, in.GetVkID())
	if err != nil {
		return &pb.UserID{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return &pb.UserID{Uid: userID}, nil
}
