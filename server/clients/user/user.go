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
	pbUserID, err := client.userClient.CreateUser(context.Background(),
		domain.ToPbUserReg(user))

	if err != nil {
		return 0, errors.Wrap(err, "user client error: ")
	}

	return pbUserID.GetUid(), nil
}

func (client *UserClient) EditUser(ctx context.Context, user domain.User) (err error) {
	_, err = client.userClient.EditUser(context.Background(),
		domain.ToPbUserEdit(user))

	if err != nil {
		if strings.Contains(err.Error(), userdomain.UserNotFoundError.Error()) {
			return domain.ErrUserNotFound
		}
		return errors.Wrap(err, "user client error: ")
	}

	return nil
}

func (client *UserClient) GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByID(context.Background(),
		&userproto.UserID{Uid: userID})

	if err != nil {
		return domain.User{}, errors.Wrap(err, "user client error: ")
	}

	return *domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	pbUser, err := client.userClient.GetUserByUsername(context.Background(),
		&userproto.Username{Username: username})

	if err != nil {
		return domain.User{}, errors.Wrap(err, "user client error: ")
	}

	return *domain.ToUser(pbUser), nil
}

func (client *UserClient) GetUsers(ctx context.Context) (users []domain.User, err error) {
	pbUsers, err := client.userClient.GetUsers(context.Background(), &userproto.Empty{})

	if err != nil {
		return nil, errors.Wrap(err, "user client error: ")
	}

	users = make([]domain.User, 0, len(pbUsers.GetUsers()))
	for _, pbUser := range pbUsers.GetUsers() {
		users = append(users, *domain.ToUser(pbUser))
	}

	return users, nil
}
