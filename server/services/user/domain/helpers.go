package domain

import (
	pb "pinterest/services/user/proto"
)

func PbUserRegToUser(pbUser *pb.UserReg) User {
	return User{
		Username:  pbUser.Username,
		Password:  pbUser.Password,
		FirstName: pbUser.FirstName,
		LastName:  pbUser.LastName,
		Email:     pbUser.Email,
	}
}

func PbUserEditinputToUser(pbUser *pb.UserEditInput) User {
	return User{
		UserID:    pbUser.UserID,
		Username:  pbUser.Username,
		FirstName: pbUser.FirstName,
		LastName:  pbUser.LastName,
		Email:     pbUser.Email,
	}
}

func UserToPbUserOutput(user User) *pb.UserOutput {
	return &pb.UserOutput{
		UserID:     user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		AvatarPath: user.AvatarPath,
	}
}

func UsersToPbUserListOutput(users []User) *pb.UsersListOutput {
	result := make([]*pb.UserOutput, 0, len(users))
	for _, user := range users {
		result = append(result, UserToPbUserOutput(user))
	}
	return &pb.UsersListOutput{
		Users: result,
	}
}
