package facade

import (
	"context"
	"pinterest/services/user/application"
	"pinterest/services/user/domain"
	pb "pinterest/services/user/proto"

	"github.com/pkg/errors"
	_ "google.golang.org/grpc"
)

type UserFacade struct {
	pb.UnimplementedUserServer
	app application.UserAppInterface
}

func NewUserFacade(app application.UserAppInterface) *UserFacade {
	return &UserFacade{
		app: app,
	}
}

func (facade *UserFacade) CreateUser(ctx context.Context, in *pb.UserReg) (*pb.UserID, error) {
	userID, err := facade.app.CreateUser(ctx, domain.PbUserRegToUser(in))
	if err != nil {
		return &pb.UserID{}, errors.Wrap(err, "Could not create user:")
	}

	return &pb.UserID{Uid: userID}, nil
}

func (facade *UserFacade) EditUser(ctx context.Context, in *pb.UserEditInput) (*pb.Empty, error) {
	err := facade.app.EditUser(ctx, domain.PbUserEditinputToUser(in))
	return &pb.Empty{}, err
}

func (facade *UserFacade) GetUserByUserID(ctx context.Context, in *pb.UserID) (*pb.UserOutput, error) {
	// TODO
	return &pb.UserOutput{}, nil
}

func (facade *UserFacade) GetUserByUsername(ctx context.Context, in *pb.Username) (*pb.UserOutput, error) {
	// TODO
	return &pb.UserOutput{}, nil
}

func (facade *UserFacade) GetUsers(ctx context.Context, in *pb.Empty) (*pb.UsersListOutput, error) {
	// TODO
	return &pb.UsersListOutput{}, nil
}

func (facade *UserFacade) ChangePassword(ctx context.Context, in *pb.Password) (*pb.Empty, error) {
	// TODO
	return &pb.Empty{}, nil
}
