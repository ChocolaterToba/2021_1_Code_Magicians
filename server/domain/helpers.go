package domain

import (
	"net/http"
	pbauth "pinterest/services/auth/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCookieInfo(pbCookieInfo *pbauth.CookieInfo, secure bool, httpOnly bool, sameSite http.SameSite) *CookieInfo {
	return &CookieInfo{
		UserID: pbCookieInfo.GetUserID(),
		Cookie: ToCookie(pbCookieInfo.GetCookie(), secure, httpOnly, sameSite),
	}
}

func ToPbCookieInfo(cookieInfo CookieInfo) *pbauth.CookieInfo {
	return &pbauth.CookieInfo{
		UserID: cookieInfo.UserID,
		Cookie: ToPbCookie(cookieInfo.Cookie),
	}
}

func ToCookie(pbCookie *pbauth.Cookie, secure bool, httpOnly bool, sameSite http.SameSite) *http.Cookie {
	return &http.Cookie{
		Name:     DefaultCookieName,
		Value:    pbCookie.GetValue(),
		Expires:  pbCookie.GetExpires().AsTime(),
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	}
}

func ToPbCookie(cookie *http.Cookie) *pbauth.Cookie {
	return &pbauth.Cookie{
		Value:   cookie.Value,
		Expires: timestamppb.New(cookie.Expires),
	}
}
