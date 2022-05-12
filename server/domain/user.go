package domain

import userpb "pinterest/services/user/proto"

type User struct {
	UserID     uint64 `json:"user_id,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	AvatarPath string `json:"avatar,omitempty"`
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

func ToUser(pbUser *userpb.UserOutput) User {
	return User{
		UserID:     uint64(pbUser.UserID),
		Username:   pbUser.GetUsername(),
		FirstName:  pbUser.GetFirstName(),
		LastName:   pbUser.GetLastName(),
		Email:      pbUser.GetEmail(),
		AvatarPath: pbUser.GetAvatarPath(),
	}
}

func ToUsers(pbUsers []*userpb.UserOutput) []User {
	result := make([]User, 0, len(pbUsers))

	for _, pbUser := range pbUsers {
		result = append(result, ToUser(pbUser))
	}

	return result
}
