package handlers

import (
	"FoodApi/src/providers"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProductFromProvider(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(providers.GetNutrition())
	json.NewEncoder(w).Encode(providers.GetNutrition())
}
