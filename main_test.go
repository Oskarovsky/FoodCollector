package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Here, we form a new HTTP request. This is the request that's going to be passed to our handler.
	req, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `Hello Food API`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v, but want %v", actual, expected)
	}
}
