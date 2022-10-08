package models

type Product struct {
	ID       int     `json:"ID"`
	Name     string  `json:"Name"`
	Energy   int     `json:"Energy"`
	Protein  float32 `json:"Protein"`
	Fat      float32 `json:"Fat"`
	Carbs    float32 `json:"Carbs"`
	Price    float32 `json:"Price"`
	Provider string  `json:"Provider"`
}
