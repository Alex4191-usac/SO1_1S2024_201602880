package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
	Id int `json:"id"`
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	
	router.GET("/data", infoHandler) 
	
	port := 8080
	router.Run(fmt.Sprintf(":%d", port))
}

func infoHandler(c *gin.Context) {

	message := Message {
		Name: "Bryan Alexander Portillo Alvarado",
		Id: 201602880,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}