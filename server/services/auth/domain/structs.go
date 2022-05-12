package domain

import "time"

type Cookie struct {
	Value   string
	Expires time.Time
}

type CookieInfo struct {
	UserID uint64
	Cookie Cookie
}
