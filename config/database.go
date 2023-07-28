package config

import (
	"github.com/abiyyu03/database-book-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper" 
)
 
var dbPass = viper.Get("DBPASS")
var dbUser = viper.Get("DBUSER")
var dbName = viper.Get("DBNAME")

var Database *gorm.DB
var DATABASE_URI string = dbUser.(string)+":"+dbPass.(string)+"@tcp(localhost:3306)/dbbook?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.Book{})

	return nil
}