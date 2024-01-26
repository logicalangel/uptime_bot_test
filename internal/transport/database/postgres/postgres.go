package postgres

import (
	"gorm.io/gorm/logger"

	"github.com/apex/log"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/transport/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	db  *gorm.DB
	cfg config.Postgres
}

func (c connection) HealthCheck() bool {
	if c.db != nil {
		return false
	}
	conn, err := c.db.DB()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	err = conn.Ping()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	return true
}

func (c connection) GetConnection() *gorm.DB {
	if c.db != nil {
		return c.db
	}

	var err error
	c.db, err = gorm.Open(postgres.Open(c.cfg.Dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	return c.db
}

func New(cfg config.Postgres) database.IDatabase {
	return &connection{
		cfg: cfg,
	}
}
