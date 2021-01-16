package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/38tter/yutabe/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
)

func quadkeyString(t maptile.Tile) string {
	s := strconv.FormatInt(int64(t.Quadkey()), 4)

	zeros := "000000000000000000000000000000"
	return zeros[:((int(t.Z)+1)-len(s))/2] + s
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	user := router.Group("/User")
	{
		user.POST("/signup", routes.UserSignUp)
	}
	fmt.Printf("connect to db")
	connectDB()

	router.Run(":8080")
}

func SaveShop() {

}

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type:varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400)"`
	Status      string `gorm:"type:char(2);not null"`
}

func connectDB() *gorm.DB {
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
	db.AutoMigrate(&Product{})
	fmt.Println("db connected: ", &db)
	return db
}
