package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"smarthome/db"
	"smarthome/handlers"
	"smarthome/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up database connection
	dbURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/smarthome")
	database, err := db.New(dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer database.Close()

	log.Println("Connected to database successfully")

	// Initialize temperature service
	temperatureAPIURL := getEnv("TEMPERATURE_API_URL", "http://temperature-api:8081")
	temperatureService := services.NewTemperatureService(temperatureAPIURL)
	log.Printf("Temperature service initialized with API URL: %s\n", temperatureAPIURL)

	// Initialize router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API routes
	apiRoutes := router.Group("/api/v1")

	// Register sensor routes
	sensorHandler := handlers.NewSensorHandler(database, temperatureService)
	sensorHandler.RegisterRoutes(apiRoutes)

	// Start server
	srv := &http.Server{
		Addr:    getEnv("PORT", ":8080"),
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exited properly")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// @title           Swagger API
// @version         1.0
// @description     This is REST API.

// @contact.name   OWNER
// @contact.url    test.ru
// @contact.email  test@test.test

// @host      localhost:8080
// @BasePath  /

// List godoc
// @Summary      List sensors
// @Description  List sensors
// @Tags         sensor
// @Accept       json
// @Produce      json
// @Param        offset   query      int  true  "Offset"
// @Param        limit   query      int  true  "Limit"
// @Success      200  {array}  Sensor
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /sensors [get]
func List() {

}

// Get godoc
// @Summary      Get sensor
// @Description  Get sensor by ID
// @Tags         sensor
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of sensor"
// @Success      200  {object}  Sensor
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /sensors/:id [get]
func Get() {

}

// Create godoc
// @Summary      Create sensor
// @Description  Create sensor
// @Tags         sensor
// @Accept       json
// @Produce      json
// @Param        data   body      object  true  "Sensor object"
// @Success      201  {object}  Sensor
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /sensors [post]
func Create() {

}

// Update godoc
// @Summary      Update sensor
// @Description  Update sensor
// @Tags         sensor
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of sensor"
// @Param        data   body      object  true  "Sensor object"
// @Success      200  {object}  Sensor
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /sensors/:id [put]
func Update() {

}

// Delete godoc
// @Summary      Delete sensor
// @Description  Delete sensor by ID
// @Tags         sensor
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of sensor"
// @Success      204
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /sensors/:id [delete]
func Delete() {

}

type Sensor struct {
	ID int
	Name string
	Location string
	Type string
	Status string
	CreatedAt string
	UpdatedAt string
}

type HTTPError struct {
	Code int
	Message string
	Fields interface{}
}