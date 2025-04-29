package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Структура для ответа с температурой
type TemperatureResponse struct {
	Location string  `json:"location"`
	sensorID string  `json:"sensorID"`
	Value    float64 `json:"value"`
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Настройка маршрутизатора Gin
	router := gin.Default()

	// Маршрут для получения температуры
	router.GET("/temperature", getTemperature)

	// Запуск сервера
	port := getEnv("PORT", "8081")
	log.Printf("Сервер запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

// Обработчик для получения температуры
func getTemperature(c *gin.Context) {
	location := c.Query("location")
	sensorID := c.Query("location")
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

	// Генерация случайной температуры от -30 до +40 градусов
	temperature := -30.0 + rand.Float64()*70.0
	temperature = float64(int(temperature*10)) / 10 // Округление до 1 знака после запятой

	// Определение статуса в зависимости от температуры

	// Формирование ответа
	response := TemperatureResponse{
		Value:    temperature,
		sensorID: sensorID,
		Location: location,
	}

	c.JSON(http.StatusOK, response)
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
