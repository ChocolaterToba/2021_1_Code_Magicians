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
		return &pb.UserID{}, errors.Wrap(err, "Could not create user")
	}

	return &pb.UserID{Uid: userID}, nil
}

func (facade *UserFacade) EditUser(ctx context.Context, in *pb.UserEditInput) (*pb.Empty, error) {
	err := facade.app.EditUser(ctx, domain.PbUserEditinputToUser(in))
	return &pb.Empty{}, err
}

func (facade *UserFacade) GetUserByID(ctx context.Context, in *pb.UserID) (*pb.UserOutput, error) {
	user, err := facade.app.GetUserByID(ctx, in.GetUid())
	if err != nil {
		return &pb.UserOutput{}, errors.Wrap(err, "Could not get user by id")
	}
	return domain.UserToPbUserOutput(user), nil
}

func (facade *UserFacade) GetUserByUsername(ctx context.Context, in *pb.Username) (*pb.UserOutput, error) {
	user, err := facade.app.GetUserByUsername(ctx, in.GetUsername())
	if err != nil {
		return &pb.UserOutput{}, errors.Wrap(err, "Could not get user by username")
	}
	return domain.UserToPbUserOutput(user), nil
}

func (facade *UserFacade) GetUsers(ctx context.Context, in *pb.Empty) (*pb.UsersListOutput, error) {
	users, err := facade.app.GetUsers(ctx)
	if err != nil {
		return &pb.UsersListOutput{}, errors.Wrap(err, "Could not get users")
	}
	return domain.UsersToPbUserListOutput(users), nil
}
