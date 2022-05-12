package facade

import (
	"bytes"
	"context"
	"io"
	"pinterest/services/user/application"
	"pinterest/services/user/domain"
	pb "pinterest/services/user/proto"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

const maxPostAvatarBodySize = 8 * 1024 * 1024 // 8 mB

func (facade *UserFacade) UpdateAvatar(stream pb.User_UpdateAvatarServer) error {

	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive user id")
	}

	userID := req.GetUserId()

	req, err = stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive filename")
	}

	filename := req.GetFilename()

	// filenamePrefix := uniuri.NewLen(10) // generating random filename
	// newAvatarPath := "avatars/" + filenamePrefix + req.GetExtension() // TODO: avatars folder sharding by date

	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}
		chunk := req.GetChunk()
		size := len(chunk)

		imageSize += size
		if imageSize > maxPostAvatarBodySize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxPostAvatarBodySize)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	err = facade.app.UpdateAvatar(context.Background(), userID, filename, &imageData)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.Empty{})
}
