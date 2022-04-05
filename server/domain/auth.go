package domain

import (
	"net/http"
	authpb "pinterest/services/auth/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserLoginInput is used when parsing JSON in auth/login handler
type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToCookieInfo(pbCookieInfo *authpb.CookieInfo, secure bool, httpOnly bool, sameSite http.SameSite) *CookieInfo {
	return &CookieInfo{
		UserID: pbCookieInfo.GetUserID(),
		Cookie: ToCookie(pbCookieInfo.GetCookie(), secure, httpOnly, sameSite),
	}
}

func ToPbCookieInfo(cookieInfo CookieInfo) *authpb.CookieInfo {
	return &authpb.CookieInfo{
		UserID: cookieInfo.UserID,
		Cookie: ToPbCookie(cookieInfo.Cookie),
	}
}

func ToCookie(pbCookie *authpb.Cookie, secure bool, httpOnly bool, sameSite http.SameSite) *http.Cookie {
	return &http.Cookie{
		Name:     DefaultCookieName,
		Value:    pbCookie.GetValue(),
		Expires:  pbCookie.GetExpires().AsTime(),
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	}
}

func ToPbCookie(cookie *http.Cookie) *authpb.Cookie {
	return &authpb.Cookie{
		Value:   cookie.Value,
		Expires: timestamppb.New(cookie.Expires),
	}
}
