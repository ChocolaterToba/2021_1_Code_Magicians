package persistence

import (
	"context"
	"fmt"
	"log"
	"pinterest/domain/entity"
	"strings"

	"github.com/jackc/pgx/v4"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepo {
	return &UserRepo{db}
}

const saveUserQuery string = "INSERT INTO Users (username, passwordhash, salt, email, first_name, last_name, avatar)\n" +
	"values ($1, $2, $3, $4, $5, $6, $7)\n" +
	"RETURNING userID"

func (r *UserRepo) SaveUser(user *entity.User) (int, error) {
	row := r.db.QueryRow(context.Background(), saveUserQuery, user.Username, user.Password, user.Salt, user.Email, user.FirstName, user.LastName, user.Avatar)
	newUserID := 0
	err := row.Scan(&newUserID)
	if err != nil {
		// If username/email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return -1, fmt.Errorf("Username or email is already taken")
		}

		// Other errors
		log.Println(err)
		return -1, err
	}
	return newUserID, nil
}

const deleteUserQuery string = "DELETE FROM Users WHERE id=$1"

func (r *UserRepo) DeleteUser(userID int) error {
	_, err := r.db.Exec(context.Background(), deleteUserQuery)
	return err
}

const getUserQuery string = "SELECT username, passwordhash, salt, email, first_name, last_name, avatar FROM Users WHERE id=$1"

func (r *UserRepo) GetUser(userID int) (*entity.User, error) {
	user := entity.User{UserID: userID}
	row := r.db.QueryRow(context.Background(), getUserQuery, userID)
	err := row.Scan(&user.Username, &user.Password, &user.Salt, &user.Email, &user.FirstName, &user.LastName, &user.Avatar)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("No user found with such id")
		}
		// Other errors
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

const getUsersQuery string = "SELECT userID, username, passwordhash, salt, email, first_name, last_name, avatar FROM Users"

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := r.db.Query(context.Background(), getUserQuery)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("No users found in database")
		}

		// Other errors
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Salt, &user.Email, &user.FirstName, &user.LastName, &user.Avatar)
		if err != nil {
			return nil, err // TODO: error handling
		}
		users = append(users, user)
	}
	return users, nil
}

const getUserByUsernameQuery string = "SELECT userID, passwordhash, salt, email, first_name, last_name, avatar\n" +
	"FROM Users WHERE username=$1"

func (r *UserRepo) GetUserByUsername(username string) (*entity.User, error) {
	user := entity.User{Username: username}
	row := r.db.QueryRow(context.Background(), getUserByUsernameQuery, username)
	err := row.Scan(&user.UserID, &user.Password, &user.Salt, &user.Email, &user.FirstName, &user.LastName, &user.Avatar)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("No user found with such username")
		}

		// Other errors
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
