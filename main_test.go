package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAPI(t *testing.T) {
	// Test case 1: Missing lat/lon parameters
	req1 := httptest.NewRequest("GET", "/api", nil)
	w1 := httptest.NewRecorder()
	handleAPI(w1, req1)
	if w1.Code != http.StatusBadRequest {
		t.Errorf("Test case 1 failed: Expected status code %d, but got %d", http.StatusBadRequest, w1.Code)
	}

	// Test case 2: Invalid lat/lon parameters
	req2 := httptest.NewRequest("GET", "/api?lat=foo&lon=bar", nil)
	w2 := httptest.NewRecorder()
	handleAPI(w2, req2)
	if w2.Code != http.StatusInternalServerError {
		t.Errorf("Test case 2 failed: Expected status code %d, but got %d", http.StatusInternalServerError, w2.Code)
	}

	// Test case 3: Successful API call
	req3 := httptest.NewRequest("GET", "/api?lat=52.52&lon=13.41", nil)
	w3 := httptest.NewRecorder()
	handleAPI(w3, req3)
	if w3.Code != http.StatusOK {
		t.Errorf("Test case 3 failed: Expected status code %d, but got %d", http.StatusOK, w3.Code)
	}
	var response1 TemperatureResponse
	err := json.NewDecoder(w3.Body).Decode(&response1)
	if err != nil {
		t.Errorf("Test case 3 failed: Error decoding JSON response")
	}
	if response1.Temp == 0 {
		t.Errorf("Test case 3 failed: Temperature value not found in response")
	}

	// Test case 4: Invalid route
	req4 := httptest.NewRequest("GET", "/foo", nil)
	w4 := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w4, req4)
	if w4.Code != http.StatusNotFound {
		t.Errorf("Test case 4 failed: Expected status code %d, but got %d", http.StatusNotFound, w4.Code)
	}
}