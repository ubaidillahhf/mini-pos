package configs

import (
	"fmt"
	"pos/models/products"
	"pos/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configuration {
	var configDB = Configuration{
		DB_USERNAME: "root",
		DB_PASSWORD: "root",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
		DB_NAME:     "pos",
	}
	return configDB
}

func InitDB() {
	configDB := GetConfig()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_USERNAME,
		configDB.DB_PASSWORD,
		configDB.DB_HOST,
		configDB.DB_PORT,
		configDB.DB_NAME)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&products.Product{})
}
