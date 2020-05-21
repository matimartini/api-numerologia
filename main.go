package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)


func main() {
	route := gin.Default()

	route.GET("/", ping)
	route.GET("/ping", ping)


	router.Run(":" + port)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context){
	io.WriteString(w, "Hello Word")
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