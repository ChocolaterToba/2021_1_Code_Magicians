package domain

import (
	pb "pinterest/services/auth/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCookie(pbCookie *pb.Cookie) Cookie {
	return Cookie{
		Value:   pbCookie.GetValue(),
		Expires: pbCookie.GetExpires().AsTime(),
	}
}

func ToPbCookie(cookie Cookie) *pb.Cookie {
	return &pb.Cookie{
		Value:   cookie.Value,
		Expires: timestamppb.New(cookie.Expires),
	}
}

func ToCookieInfo(pbCookieInfo *pb.CookieInfo) CookieInfo {
	return CookieInfo{
		UserID: pbCookieInfo.GetUserID(),
		Cookie: ToCookie(pbCookieInfo.GetCookie()),
	}
}

func ToPbCookieInfo(cookieInfo CookieInfo) *pb.CookieInfo {
	return &pb.CookieInfo{
		UserID: cookieInfo.UserID,
		Cookie: ToPbCookie(cookieInfo.Cookie),
	}
}
