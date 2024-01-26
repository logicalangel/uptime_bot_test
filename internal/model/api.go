package model

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"gorm.io/gorm"
)

type ApiStatus string
type ApiMethod string

const (
	ApiRunningStatus ApiStatus = "RUNNING"
	ApiStopedStatus  ApiStatus = "STOPED"
)

const (
	GetMethod    ApiMethod = "GET"
	PostMethod   ApiMethod = "POST"
	PutMethod    ApiMethod = "PUT"
	DeleteMethod ApiMethod = "DELETE"
)

type Api struct {
	gorm.Model

	Status       ApiStatus       `json:"status"`
	Url          string          `json:"url"`
	Method       ApiMethod       `json:"method"`
	CallInterval uint            `json:"call_interval"`
	Body         string          `json:"body"`
	Headers      gormjsonb.JSONB `gorm:"type:jsonb;default:'[]';not null" json:"headers"`
	Webhook      string          `json:"webhook"`
}
