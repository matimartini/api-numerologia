package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matimartini/api-numerologia/controller"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Println("$PORT must be set")
	}

	router := gin.New()

	router.GET("/", hello)
	router.GET("/ping", ping)
	router.GET("calculate", getCalculate)

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
	controller := controller.PersonController{
		Context: c,
	}
	controller.GetPerson()

	/*id := c.Query("name")
	fecha := c.Query("fecha")

	log.Println("id send param:", id)

	p := &service.Person{}
	p.NewPerson(id, fecha)
	c.JSON(200, gin.H{
		"message": "Num: " + p.FullName + " Fecha: " + p.FullName,
	})*/
}
