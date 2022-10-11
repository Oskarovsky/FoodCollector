package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Identifier int `json:"fdcId"`
}

func GetNutrition() string {
	response, err := http.Get("https://api.nal.usda.gov/fdc/v1/food/2057648?api_key=JU9ViL2Rf0dYov0mXG82QLUeAU0bUEQrI1KfAMBF")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Fill the record with the data from the JSON
	var record Product

	json.Unmarshal(responseData, &record)
	fmt.Printf("API Response as struct %+v\n", record)

	return string(responseData)
}
