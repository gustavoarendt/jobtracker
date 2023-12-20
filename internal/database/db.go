package database

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/gustavoarendt/jobtracker/configs"
	"github.com/gustavoarendt/jobtracker/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	err          error
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
)

func DbConnection(config *configs.Config) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&entities.User{}, &entities.Company{}, &entities.Job{})
}

func JwtConnection(config *configs.Config) {
	Jwt = config.TokenAuth
	JwtExpiresIn = config.JWTExpiresIn
}
