package usecase

import (
	"fmt"
	"net/http"

	"github.com/38tter/yutabe/infrastructure"
	"github.com/38tter/yutabe/models"
	"github.com/gin-gonic/gin"
)

type ProductResponse struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Memo        string `json:"memo"`
	Status      string `json:"status"`
}

func (res *ProductResponse) serialize(rawProduct *models.Product) {
	res.ID = rawProduct.ID
	res.ProductName = rawProduct.ProductName
	res.Memo = rawProduct.Memo
	res.Status = rawProduct.Status
}

func GetProducts(ctx *gin.Context) {
	db := infrastructure.ConnectDB()
	var products []models.Product
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
}
