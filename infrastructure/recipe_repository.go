package infrastructure

import "github.com/38tter/yutabe/models"

func SaveRecipe(r *models.Recipe) {
	db := ConnectDB()
	db.Create(&r)
	defer db.Close()
}
