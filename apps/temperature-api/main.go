package main

import (
	   "math/rand"
	   "net/http"
	   "time"
	   "encoding/json"
)

type TemperatureResponse struct {
	Location  string  `json:"location"`
	SensorID  string  `json:"sensorId"`
	Temperature float64 `json:"temperature"`
}

func getTemperature(location, sensorID string) (string, string) {
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
	return location, sensorID
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")
	location, sensorID = getTemperature(location, sensorID)

	rand.Seed(time.Now().UnixNano())
	temperature := 15.0 + rand.Float64()*10.0 // 15-25°C

	resp := TemperatureResponse{
		Location: location,
		SensorID: sensorID,
		Temperature: temperature,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/temperature", temperatureHandler)
	http.ListenAndServe(":8081", nil)
}
