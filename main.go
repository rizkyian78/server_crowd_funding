package main

import (
	"crowd_fund_server/Users"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "rizkyian78:Rizkyian_78@tcp(127.0.0.1:3306)/CrowdFunding_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := Users.NewRepository(db)
	user := Users.User{
		ID:   uuid.NewString(),
		Name: "Testing",
	}
	userRepository.Save(user)
}
