package database

import (
	"chat_server/adapter/database/model"
	"chat_server/config"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {

	gormConfig := &gorm.Config{}

	db, err := gorm.Open(postgres.Open(config.DSN()), gormConfig)
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return errors.New("failed to migrate database")
	}

	return nil
}