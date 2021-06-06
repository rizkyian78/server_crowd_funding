package main

import (
	"crowd_fund_server/Users"
	"crowd_fund_server/handler"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	port_db := os.Getenv("PORT_DB")
	DB_NAME := os.Getenv("DB_NAME")
	port_host := os.Getenv("PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, port_db, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := Users.NewRepository(db)
	userService := Users.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.New()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(fmt.Sprintf(":%s", port_host))
}
