package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port = os.Getenv("PORT")

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	data, err := os.ReadFile("data/data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json: %v", err)
	}

	var jsonData []interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	r.GET("/videos/*filepath", func(c *gin.Context) {
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		c.File("data/videos" + c.Param("filepath"))
	})

	r.Static("/images", "data/images")
	r.Static("/headers", "data/headers")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, jsonData)
	})

	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
