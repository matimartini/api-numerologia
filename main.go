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

	router.GET("/", ping)
	router.GET("/ping", ping)

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

// function without  gin-gonic
/*func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Word")
	fmt.Println("Hello Word!!!!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}*/
