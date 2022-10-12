package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Identifier int    `json:"fdcId"`
	Name       string `json:"description"`
}

func GetNutritionByProductId(productId int) string {
	req, err := http.NewRequest(http.MethodGet, "https://api.nal.usda.gov/fdc/v1/food/"+strconv.Itoa(productId), nil)

	q := req.URL.Query()
	q.Add("api_key", "JU9ViL2Rf0dYov0mXG82QLUeAU0bUEQrI1KfAMBF")
	req.URL.RawQuery = q.Encode()

	response, err := http.Get(req.URL.String())

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

func GetNutritionByProductName(productName string) string {

	var jsonData = []byte(`{
		"query": "` + productName + `",
		"pageSize": 1
	}`)

	fmt.Printf("TEST: %d", jsonData)
	req, err := http.NewRequest(http.MethodPost, "https://api.nal.usda.gov/fdc/v1/foods/search", bytes.NewBuffer(jsonData))

	q := req.URL.Query()
	q.Add("api_key", "JU9ViL2Rf0dYov0mXG82QLUeAU0bUEQrI1KfAMBF")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(req)
	//response, err := http.Post(req)

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
