package main

import (
	"bwa-golang/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("connected to database")

	userRepository := user.NewRepository(db)

	user := user.User{
		Name: "Test Nama Saya",
	}

	userRepository.Save(user)

}
