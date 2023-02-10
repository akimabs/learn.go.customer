package lib

import (
	"customer/internal/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() Database {
	var (
		connection *gorm.DB
		err        error
	)

	godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL"),
	)

	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	connection.AutoMigrate(&models.Companies{}, &models.Roles{}, &models.Users{}, &models.Devices{}, &models.Features{})

	if err != nil {
		fmt.Println(err)
	}

	return Database{
		DB: connection,
	}
}
