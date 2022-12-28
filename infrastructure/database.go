package infrastructure

import (
	"bytes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"ojire/config"
	"strconv"
	"time"
)

type HandlerDatabase struct {
	DatabaseClient *gorm.DB
}

func NewDatabaseClient(e *config.Config) (*HandlerDatabase, error) {
	databaseInit := &HandlerDatabase{}
	slave, err := databaseInit.OpenPostgresConnectionRead(e)
	if err != nil {
		return nil, err
	}
	databaseInit.DatabaseClient = slave
	return databaseInit, nil
}

func (hd *HandlerDatabase) OpenPostgresConnectionRead(e *config.Config) (*gorm.DB, error) {
	var bufferSlave bytes.Buffer
	bufferSlave.WriteString("host=" + e.Database.Host)
	bufferSlave.WriteString(" port=" + strconv.Itoa(e.Database.Port))
	bufferSlave.WriteString(" dbname=" + e.Database.DBName)
	bufferSlave.WriteString(" sslmode=disable")
	connectionSlave := bufferSlave.String()
	dbRead, err := gorm.Open("postgres", connectionSlave)
	if err != nil {
		return nil, err
	}
	dbRead.LogMode(e.Database.Debug)
	dbRead.DB().SetMaxIdleConns(e.Database.SetMaxIdleCons)
	dbRead.DB().SetMaxOpenConns(e.Database.SetMaxOpenCons)
	dbRead.DB().SetConnMaxIdleTime(time.Duration(e.Database.SetConMaxIdleTime) * time.Minute)
	dbRead.DB().SetConnMaxLifetime(time.Duration(e.Database.SetConMaxLifetime) * time.Minute)

	return dbRead, nil
}
