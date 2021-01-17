package models

type NutritionInt int

type Nutrition struct {
	ID            NutritionInt
	Fat           NutritionInt
	Protein       NutritionInt
	Sugar         NutritionInt
	Fiber         NutritionInt
	Calories      NutritionInt
	Carbohydrates NutritionInt
}
