package database

import (
	"fmt"
	"gochatserver/app/config"
	"gochatserver/app/database/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
		config.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteMigrations(db *gorm.DB) {
	// add entities to migrate
	log.Println("Executando migrations!")
	db.AutoMigrate(entities.User{})
}
