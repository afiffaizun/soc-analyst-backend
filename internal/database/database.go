package database

import (
	"fmt"
	"log"

	"github.com/afiffaizun/soc-analyst-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(host, user, password, dbname string, port int) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
    var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Connected to database")

	//Auto migrate models
	err = DB.AutoMigrate(&models.Service{}, &models.Team{}, &models.Article{})

	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}
}