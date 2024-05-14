package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AutoApproveLog struct {
	gorm.Model
	UserId    int64
	FirstName string
	ChatId    int64
	ChatTitle string
}

var DB *gorm.DB

func migrateDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("<< Failed to connect database . Exiting Now >>")
	}
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = DB.AutoMigrate(&AutoApproveLog{})
	if err != nil {
		panic("<< Failed to Automigrate Models >>")
	}

}
