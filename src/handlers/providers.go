package handlers

import (
	"FoodApi/src/providers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetProductFromProviderByProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	productId, _ := strconv.Atoi(mux.Vars(r)["productId"])
	fmt.Println(providers.GetNutritionByProductId(productId))
	json.NewEncoder(w).Encode(providers.GetNutritionByProductId(productId))
}

func GetProductFromProviderByProductName(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	productName, _ := mux.Vars(r)["productName"]
	fmt.Println(providers.GetNutritionByProductName(productName))
	json.NewEncoder(w).Encode(providers.GetNutritionByProductName(productName))
}
