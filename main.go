package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guptahardik/teradesi-news/run_api"
)

func main() {
	fmt.Println("Hello, World!")
	run_api.Give_news("bitcoin", "us", "business")
	router_business()
}

func router_business() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "news",
		})
	})
	r.POST("/news", func(c *gin.Context) {
		var data struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.BindJSON(&data)
		c.JSON(200, gin.H{"Welcome": data.Name, "Your age is Age": data.Age})
	})
	r.Run()
}
