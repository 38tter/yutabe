package infrastructure

import (
	"fmt"

	"github.com/38tter/yutabe/models"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "yutabe"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	db.AutoMigrate(&models.Product{})
	fmt.Println("db connected: ", &db)
	return db
}
