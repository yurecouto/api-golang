package database

import (
	"api-golang/src/config"
	"api-golang/src/models"
	"api-golang/src/utils"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	var erro error
	dsn := config.ConnectString

	db, erro = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if erro != nil {
		utils.Error(erro)

		return db, erro
	}

	db.AutoMigrate(&models.User{}, &models.UserToken{})

	return db, nil
}
