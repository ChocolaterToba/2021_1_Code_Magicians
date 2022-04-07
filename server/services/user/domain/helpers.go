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
