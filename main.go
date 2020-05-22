package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()

	router.GET("/", hello)
	router.GET("/ping", ping)
	router.GET("calculate/:id")

	router.Run(":" + port)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Word!!!",
	})
}

func getCalculate(c *gin.Context) {
	id := c.Param("id")

	log.Println("id send param:", id)
	c.JSON(200, gin.H{
		"message": "Num: " + id,
	})
}
