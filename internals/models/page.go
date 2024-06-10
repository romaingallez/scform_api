package models

import (
	"github.com/go-rod/rod"
)

type Page struct {
	Page *rod.Page
	User User
}
