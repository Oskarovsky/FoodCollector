package mocks

import (
	"FoodApi/src/models"
)

var Products = []models.Product{
	{
		ID:       1,
		Name:     "Apple",
		Energy:   310,
		Protein:  2.10,
		Fat:      0.01,
		Carbs:    1.99,
		Price:    0.99,
		Provider: "Tesco",
	},
	{
		ID:       2,
		Name:     "Orange",
		Energy:   37,
		Protein:  0.15,
		Fat:      0.05,
		Carbs:    4.06,
		Price:    1.29,
		Provider: "Tesco",
	},
}
