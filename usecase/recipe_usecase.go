package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/38tter/yutabe/infrastructure"
	"github.com/38tter/yutabe/models"
	"github.com/gin-gonic/gin"
)

func SaveRecipes(ctx *gin.Context) {
	var recipe models.RecipeJson
	body := ctx.Request.Body
	if err := json.NewDecoder(body).Decode(&recipe); err != nil {
		fmt.Printf("failed to save recipes: %s", err)
	}
	fmt.Printf("recipe = %#v \n", recipe)
	fmt.Printf("instruction = %#v \n", recipe.Instruction)

	var texts []models.InstructionText
	for _, t := range recipe.Instruction.Texts {
		texts = append(texts, models.InstructionText{Text: t})
	}
	r := models.Recipe{
		Name:        recipe.Name,
		ImageUrl:    recipe.ImageUrl,
		Description: recipe.Description,
		Instruction: &models.Instruction{
			Texts: texts,
		},
	}
	infrastructure.SaveRecipe(&r)

	ctx.JSON(http.StatusCreated, "")
}
