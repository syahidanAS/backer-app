package main

import (
	"backer-app/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Koneksi Database
	dsn := "root:@tcp(127.0.0.1:3306)/backer-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Tes Simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "Programmer"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
