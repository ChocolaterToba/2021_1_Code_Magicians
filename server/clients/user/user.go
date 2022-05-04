package user

import (
	"context"
	"pinterest/domain"
	userdomain "pinterest/services/user/domain"
	userproto "pinterest/services/user/proto"
	"strings"

	"github.com/pkg/errors"
)

type UserClientInterface interface {
	CreateUser(ctx context.Context, user domain.User) (userID uint64, err error)
	EditUser(ctx context.Context, user domain.User) (err error)
	GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error)
	GetUserByUsername(ctx context.Context, username string) (user domain.User, err error)
	GetUsers(ctx context.Context) (users []domain.User, err error)
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
		return 0, errors.Wrap(err, "user client error: ")
	}

	return pbUserID.GetUid(), nil
}

func (client *UserClient) EditUser(ctx context.Context, user domain.User) (err error) {
	_, err = client.userClient.EditUser(ctx, domain.ToPbUserEdit(user))

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.ErrUserNotFound
		}
		return errors.Wrap(err, "user client error: ")
	}

	return nil
}

func (client *UserClient) GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByID(ctx, &userproto.UserID{Uid: userID})

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, errors.Wrap(err, "user client error: ")
	}

	return domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByUsername(ctx, &userproto.Username{Username: username})

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, errors.Wrap(err, "user client error: ")
	}

	return domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUsers(ctx context.Context) (users []domain.User, err error) {
	pbUsers, err := client.userClient.GetUsers(ctx, &userproto.Empty{})

	if err != nil {
		return nil, errors.Wrap(err, "user client error: ")
	}

	return domain.ToUsers(pbUsers.Users), nil
}
