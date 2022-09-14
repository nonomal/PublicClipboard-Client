package model

import (
	"gorm.io/gorm"
)

type Clipboard struct {
	gorm.Model
	Msg       string `json:"msg"`
	DeleteTag bool
}

type Result struct {
	Code      int       `json:"code"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Clipboard Clipboard `json:"clipboard"`
}
