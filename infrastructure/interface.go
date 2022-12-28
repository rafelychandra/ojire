package infrastructure

import (
	"github.com/jinzhu/gorm"
	"ojire/config"
)

type InterfaceMongoDB interface {
	OpenPostgresConnectionRead(e *config.Config) (*gorm.DB, error)
}
