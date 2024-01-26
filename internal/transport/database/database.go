package database

import "gorm.io/gorm"

type IDatabase interface {
	GetConnection() *gorm.DB
	HealthCheck() bool
}
