package database

import (
	"fmt"
	
	"discipline/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(host, user, password, dbname string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate les modèles
	err = db.AutoMigrate(
		&models.User{},
		&models.Match{},
		&models.Round{},
		&models.Action{},
		&models.Report{},
		&models.Objective{},
	)
	
	return db, err
} 