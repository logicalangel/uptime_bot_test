package model

import (
	"gorm.io/gorm"
)

type Call struct {
	gorm.Model

	ApiId uint `json:"api_id"`

	Status uint `json:"status"`
	Time   uint `json:"time"`
}
