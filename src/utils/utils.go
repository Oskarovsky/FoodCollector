package utils

import "net/http"

// The ReturnJsonResponse function will convert any response
// that the HTTP server returns and set the Content-type as JSON
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}
