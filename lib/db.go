package lib

import (
	"fmt"
	"github.com/binganao/dailyCVE/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	database := viper.Get("mysql.database")
	databaseUser := viper.Get("mysql.username")
	databasePwd := viper.Get("mysql.password")
	charset := viper.Get("mysql.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		databaseUser,
		databasePwd,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.CVE{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
