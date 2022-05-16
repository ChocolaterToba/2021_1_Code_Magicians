package user

import (
	"bufio"
	"context"
	"io"
	"pinterest/domain"
	userdomain "pinterest/services/user/domain"
	userproto "pinterest/services/user/proto"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type UserClientInterface interface {
	CreateUser(ctx context.Context, user domain.User) (userID uint64, err error)
	EditUser(ctx context.Context, user domain.User) (err error)
	GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error)
	GetUserByUsername(ctx context.Context, username string) (user domain.User, err error)
	GetUsers(ctx context.Context) (users []domain.User, err error)
	UpdateAvatar(ctx context.Context, userID uint64, filename string, file io.Reader) (err error)
	GetRoles(ctx context.Context, userID uint64) (roles []string, err error)
	AddRole(ctx context.Context, userID uint64, role string) (err error)
}

type UserClient struct {
	userClient userproto.UserClient
}

func NewUserClient(userClient userproto.UserClient) *UserClient {
	return &UserClient{
		userClient: userClient,
	}
}

func (client *UserClient) CreateUser(ctx context.Context, user domain.User) (userID uint64, err error) {
	pbUserID, err := client.userClient.CreateUser(ctx, domain.ToPbUserReg(user))

	if err != nil {
		return 0, errors.Wrap(err, "user client error")
	}

	return pbUserID.GetUid(), nil
}

func (client *UserClient) EditUser(ctx context.Context, user domain.User) (err error) {
	_, err = client.userClient.EditUser(ctx, domain.ToPbUserEdit(user))

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.ErrUserNotFound
		}
		return errors.Wrap(err, "user client error")
	}

	return nil
}

func (client *UserClient) GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByID(ctx, &userproto.UserID{Uid: userID})

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, errors.Wrap(err, "user client error")
	}

	return domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByUsername(ctx, &userproto.Username{Username: username})

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, errors.Wrap(err, "user client error")
	}

	return domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUsers(ctx context.Context) (users []domain.User, err error) {
	pbUsers, err := client.userClient.GetUsers(ctx, &userproto.Empty{})

	if err != nil {
		return nil, errors.Wrap(err, "user client error")
	}

	return domain.ToUsers(pbUsers.Users), nil
}

func (client *UserClient) UpdateAvatar(ctx context.Context, userID uint64, filename string, file io.Reader) (err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := client.userClient.UpdateAvatar(ctx)
	if err != nil {
		return errors.Wrap(err, "Cannot start stream")
	}

	userIdReq := &userproto.UpdateAvatarRequest{
		Data: &userproto.UpdateAvatarRequest_UserId{
			UserId: userID,
		},
	}
	err = stream.Send(userIdReq)
	if err != nil {
		return errors.Wrap(err, "Cannot send user id to service")
	}

	filenameReq := &userproto.UpdateAvatarRequest{
		Data: &userproto.UpdateAvatarRequest_Filename{
			Filename: filename,
		},
	}
	err = stream.Send(filenameReq)
	if err != nil {
		return errors.Wrap(err, "Cannot send image's filename to service")
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

		req := &userproto.UpdateAvatarRequest{
			Data: &userproto.UpdateAvatarRequest_Chunk{
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

func (client *UserClient) GetRoles(ctx context.Context, userID uint64) (roles []string, err error) {
	pbRoles, err := client.userClient.GetRoles(ctx, &userproto.UserID{Uid: userID})

	if err != nil {
		return nil, errors.Wrap(err, "user client error")
	}

	return domain.ToRoles(pbRoles.Roles), nil
}

func (client *UserClient) AddRole(ctx context.Context, userID uint64, role string) (err error) {
	_, err = client.userClient.AddRole(ctx, &userproto.AddRoleRequest{
		UserId: userID,
		Role:   userproto.Role(userproto.Role_value[role]),
	})

	if err != nil {
		return errors.Wrap(err, "user client error")
	}

	return nil
}
