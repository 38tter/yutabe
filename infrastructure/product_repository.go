package infrastructure

import "github.com/38tter/yutabe/models"

func SaveProduct(p *models.Product) {
	db := ConnectDB()
	db.Create(&p)
	defer db.Close()
}
