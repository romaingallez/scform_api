package models

import "github.com/go-rod/rod/lib/proto"

type User struct {
	ID       int                    `json:"id"`
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Cookies  []*proto.NetworkCookie `json:"cookies"`
}
