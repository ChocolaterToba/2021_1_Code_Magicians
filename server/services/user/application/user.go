package application

import (
	"context"
	"pinterest/services/user/domain"
	repository "pinterest/services/user/infrastructure"

	"golang.org/x/crypto/bcrypt"
)

type UserAppInterface interface {
	CreateUser(ctx context.Context, user domain.User) (userID uint64, err error)
	GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error)
	GetUserByUsername(ctx context.Context, username string) (user domain.User, err error)
	GetUsers(ctx context.Context) (users []domain.User, err error)
	EditUser(ctx context.Context, user domain.User) (err error)
}

type UserApp struct {
	repo repository.UserRepoInterface
}

func NewUserApp(repo repository.UserRepoInterface) *UserApp {
	return &UserApp{
		repo: repo,
	}
}

func (app *UserApp) CreateUser(ctx context.Context, user domain.User) (userID uint64, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	return app.repo.CreateUser(ctx, user, passwordHash)
}

func (app *UserApp) GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error) {
	return app.repo.GetUserByID(ctx, userID)
}

func (app *UserApp) GetUserByUsername(ctx context.Context, username string) (user domain.User, err error) {
	return app.repo.GetUserByUsername(ctx, username)
}

func (app *UserApp) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return app.repo.GetUsers(ctx)
}

func (app *UserApp) EditUser(ctx context.Context, user domain.User) (err error) {
	//TODO: add transactions here?
	dbUser, err := app.repo.GetUserByID(ctx, user.UserID)
	if err != nil {
		return err
	}

	if user.Email != "" {
		dbUser.Email = user.Email
	}
	if user.FirstName != "" {
		dbUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		dbUser.LastName = user.LastName
	}

	return app.repo.UpdateUser(ctx, dbUser)
}
