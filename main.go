package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type TemperatureResponse struct {
	Temp float64 `json:"temp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api", handleAPI)
	http.ListenAndServe(":"+port, nil)
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" || lon == "" {
		http.Error(w, "Missing lat/lon parameters", http.StatusBadRequest)
		return
	}

	temp, err := getTemp(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := TemperatureResponse{Temp: temp}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getTemp(lat, lon string) (float64, error) {
	queryValues := url.Values{}
	queryValues.Set("latitude", lat)
	queryValues.Set("longitude", lon)
	queryValues.Set("current_weather", "true")

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", queryValues.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Open Meteo API returned HTTP %d", resp.StatusCode)
	}

	currentWeather, ok := data["current_weather"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Unable to get current weather from API response")
	}

	temp, ok := currentWeather["temperature"].(float64)
	if !ok {
		return 0, fmt.Errorf("Unable to get temperature from API response")
	}

	return temp, nil
}

//https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m