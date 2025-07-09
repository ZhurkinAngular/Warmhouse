package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TemperatureResponse struct {
	Temperature float64   `json:"temperature"`
	Value       float64   `json:"value"`
	Location    string    `json:"location"`
	SensorID    string    `json:"sensorId"`
	Timestamp   time.Time `json:"timestamp"`
	Unit        string    `json:"unit"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}

func main() {
	r := gin.Default()

	r.GET("/temperature", func(c *gin.Context) {
		location := c.Query("location")
		sensorID := c.Query("sensorId")

		// If no location is provided, use a default based on sensor ID
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

		// If no sensor ID is provided, generate one based on location
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

		// Generate random temperature between 18.0 and 25.0
		temperature := 18.0 + rand.Float64()*7.0

		response := TemperatureResponse{
			Temperature: temperature,
			Value:       temperature,
			Location:    location,
			SensorID:    sensorID,
			Timestamp:   time.Now(),
			Unit:        "°C",
			Status:      "active",
			Description: fmt.Sprintf("Temperature sensor in %s", location),
		}

		c.JSON(http.StatusOK, response)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	fmt.Println("Temperature API starting on port 8081...")
	log.Fatal(r.Run(":8081"))
}
