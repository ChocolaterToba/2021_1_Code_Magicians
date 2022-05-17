package application

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	s3client "pinterest/services/user/clients/s3"
	"pinterest/services/user/domain"
	repository "pinterest/services/user/infrastructure"
	"time"

	"github.com/dchest/uniuri"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserAppInterface interface {
	CreateUser(ctx context.Context, user domain.User) (userID uint64, err error)
	GetUserByID(ctx context.Context, userID uint64) (user domain.User, err error)
	GetUserByUsername(ctx context.Context, username string) (user domain.User, err error)
	GetUsers(ctx context.Context) (users []domain.User, err error)
	EditUser(ctx context.Context, user domain.User) (err error)
	UpdateAvatar(ctx context.Context, userID uint64, filename string, file *bytes.Buffer) (err error)
	GetRoles(ctx context.Context, userID uint64) (roles []string, err error)
	AddRole(ctx context.Context, userID uint64, role string) (err error)
}

type UserApp struct {
	repo     repository.UserRepoInterface
	s3Client s3client.S3ClientInterface
}

func NewUserApp(repo repository.UserRepoInterface, s3Client s3client.S3ClientInterface) *UserApp {
	return &UserApp{
		repo:     repo,
		s3Client: s3Client,
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

const AvatarIDLen = 10

func (app *UserApp) UpdateAvatar(ctx context.Context, userID uint64, filename string, file *bytes.Buffer) (err error) {
	user, err := app.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	extension := filepath.Ext(filename)
	if extension != ".jpeg" && extension != ".jpg" && extension != ".png" {
		return errors.Wrap(domain.UnsupportedExtensionError, fmt.Sprintf("Extension: %s", extension))
	}

	filePrefix := time.Now().Format("2006/01/02") // is used for easy manual file search

	fileID := uniuri.NewLen(AvatarIDLen)

	newAvatarPath := filePrefix + "/" + fileID + "_" + filename

	err = app.s3Client.UploadFile(ctx, newAvatarPath, file)
	if err != nil {
		return err
	}

	oldAvatarPath := user.AvatarPath
	user.AvatarPath = newAvatarPath

	err = app.repo.UpdateUser(ctx, user)
	if err != nil {
		app.s3Client.DeleteFile(ctx, newAvatarPath) // Try to delete freshly uploaded file
		return err
	}

	if oldAvatarPath == "" {
		return nil
	}

	return app.s3Client.DeleteFile(ctx, oldAvatarPath)
}

func (app *UserApp) GetRoles(ctx context.Context, userID uint64) (roles []string, err error) {
	return app.repo.GetRoles(ctx, userID)
}

func (app *UserApp) AddRole(ctx context.Context, userID uint64, role string) (err error) {
	roles, err := app.repo.GetRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, dbRole := range roles {
		if dbRole == role {
			return nil // skip role assignment
		}
	}
	return app.repo.UpdateRoles(ctx, userID, append(roles, role))
}
