package facade

import (
	"context"
	"pinterest/services/auth/application"
	"pinterest/services/auth/domain"
	pb "pinterest/services/auth/proto"

	"github.com/pkg/errors"
	_ "google.golang.org/grpc"
)

type AuthFacade struct {
	app application.AuthAppInterface
}

func NewAuthFacade(app application.AuthAppInterface) *AuthFacade {
	return &AuthFacade{
		app: app,
	}
}

func (facade *AuthFacade) LoginUser(ctx context.Context, in *pb.UserAuth) (*pb.CookieInfo, error) {
	cookieInfo, err := facade.app.LoginUser(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return &pb.CookieInfo{}, errors.Wrap(err, "Could not login user credentials:")
	}

	return domain.ToPbCookieInfo(cookieInfo), nil
}

func (facade *AuthFacade) SearchCookieByValue(ctx context.Context, in *pb.CookieValue) (*pb.CookieInfo, error) {
	result, err := facade.app.SearchCookieByValue(ctx, in.GetCookieValue())
	if err != nil {
		return &pb.CookieInfo{}, errors.Wrap(err, "Could not find cookie by value:")
	}

	return domain.ToPbCookieInfo(result), nil
}

func (facade *AuthFacade) SearchCookieByUserID(ctx context.Context, in *pb.UserID) (*pb.CookieInfo, error) {
	result, err := facade.app.SearchCookieByUserID(ctx, in.GetUid())
	if err != nil {
		return &pb.CookieInfo{}, errors.Wrap(err, "Could not find cookie by user id:")
	}

	return domain.ToPbCookieInfo(result), nil
}

func (facade *AuthFacade) LogoutUser(ctx context.Context, in *pb.CookieValue) (*pb.Error, error) {
	err := facade.app.LogoutUser(ctx, in.GetCookieValue())
	if err != nil {
		return &pb.Error{}, errors.Wrap(err, "Could not check user credentials:")
	}

	return &pb.Error{}, nil
}
