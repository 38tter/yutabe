package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/38tter/yutabe/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
)

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type:varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400)"`
	Status      string `gorm:"type:char(2);not null"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Memo        string `json:"memo"`
	Status      string `json:"status"`
}

func (res *ProductResponse) serialize(rawProduct *Product) {
	res.ID = rawProduct.ID
	res.ProductName = rawProduct.ProductName
	res.Memo = rawProduct.Memo
	res.Status = rawProduct.Status
}

func quadkeyString(t maptile.Tile) string {
	s := strconv.FormatInt(int64(t.Quadkey()), 4)

	zeros := "000000000000000000000000000000"
	return zeros[:((int(t.Z)+1)-len(s))/2] + s
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/products", func(ctx *gin.Context) {
		db := connectDB()
		var products []Product
		res := db.Find(&products)
		if res.Error != nil {
			fmt.Errorf("Error occured when all products ")
		}
		var productResponse []ProductResponse
		for _, p := range products {
			var rd ProductResponse
			rd.serialize(&p)
			productResponse = append(productResponse, rd)
		}
		ctx.JSON(http.StatusOK, productResponse)
	})

	user := router.Group("/User")
	{
		user.POST("/signup", routes.UserSignUp)
	}
	fmt.Printf("connect to db")
	connectDB()

	_ = Product{
		ProductName: "ラーメン",
		Memo:        "ひごもんずのラーメンは熊本風。最高",
		Status:      "CREATED",
	}
	//saveProduct(&p)

	router.Run(":8081")
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
