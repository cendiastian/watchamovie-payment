package driver

import (
	"fmt"
	"log"
	"watchamovie-payment/config"
	inData "watchamovie-payment/features/invoice/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println("Config : ", config)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db

	DB.AutoMigrate(&inData.Invoice{})
}
