package infrastructure

import (
	"fmt"

	"github.com/38tter/yutabe/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db.AutoMigrate(&models.Ingredient{})
	db.AutoMigrate(&models.Instruction{})
	db.AutoMigrate(&models.InstructionText{})
	db.AutoMigrate(&models.Recipe{})
	fmt.Println("db connected: ", &db)
	return db
}
