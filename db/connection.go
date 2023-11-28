package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DatabaseConn struct{}

//var DB *gorm.DB

func (d DatabaseConn) GetDB() *gorm.DB {
	//Connection Database
	dsn := fmt.Sprintf(`%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, viper.GetString("DBUSER"), viper.GetString("DBHOST"), viper.GetString("DBPORT"), viper.GetString("DBNAME"))
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatal("Database Connection Failed", err)
	}

	//Migration Database
	Migrate(DB)

	return DB
}
