package domain

import userpb "pinterest/services/user/proto"

type User struct {
	UserID    uint64 `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserIDResponse struct {
	UserID uint64 `json:"user_id"`
}

func ToPbUserReg(user User) *userpb.UserReg {
	return &userpb.UserReg{
		Username:  user.Username,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToPbUserEdit(user User) *userpb.UserEditInput {
	return &userpb.UserEditInput{
		UserID:    user.UserID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToUser(pbuser *userpb.UserOutput) *User {
	return &User{
		UserID:    uint64(pbuser.UserID),
		Username:  pbuser.GetUsername(),
		FirstName: pbuser.GetFirstName(),
		LastName:  pbuser.GetLastName(),
		Email:     pbuser.GetEmail(),
	}
}
