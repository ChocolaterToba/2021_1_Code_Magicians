package domain

import (
	"net/http"
)

const (
	DefaultCookieName = "session_id"
)

// CookieInfo contains information about a cookie: which user it belongs to and cookie itself
type CookieInfo struct {
	UserID uint64
	Cookie *http.Cookie
}
