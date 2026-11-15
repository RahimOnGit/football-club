package database

import (
	"fmt"
	"football-club/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	var err error
	DB , err = gorm.Open(sqlite.Open("football.db"), &gorm.Config{})
	if err !=nil{
		panic("failed to connect database")
	}

//	auto mig

	err = DB.AutoMigrate(&models.Player{})
	if err != nil {
		panic("failed to migrate database")
	}

//	seed

 var count int64
DB.Model(&models.Player{}).Count(&count)
	if count == 0 {
		fmt.Print("seeding the db")
		DB.Create(&models.Player{Name: "Semo" , Position: "Mid" , Club: "Elsababi" , Goals: 4})
		DB.Create(&models.Player{Name: "Mo" , Position: "Goalkeeper" , Club: "Elsababi" , Goals: 0})
		fmt.Print("the db seeded with players")
	}


}