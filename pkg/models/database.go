package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Gorm *gorm.DB
}

func ConnectDB() (*Database, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	// 環境変数からセッションキーを取得する
	dbname := os.Getenv("DATABASE_NAME")
	if dbname == "" {
		log.Fatal("SESSION_KEY is required")
	}

	gormDB, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	gormDB.AutoMigrate(&User{})
	return &Database{Gorm: gormDB}, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.Gorm.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
