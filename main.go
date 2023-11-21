package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type periferico struct {
	ID    string  `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

var perifericos = []periferico{
	{ID: "789", Nome: "Hyperx Quadcast", Preco: 999.88},
}

func main() {
	router := gin.Default()
	router.GET("/perifericos", getPerifericos)
	router.POST("/perifericos", postPerifericos)

	router.Run("localhost:8080")
}

func getPerifericos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, perifericos)
}

func postPerifericos(c *gin.Context) {
	var newPerifericos []periferico

	if err := c.BindJSON(&newPerifericos); err != nil {
		return
	}

	perifericos = append(perifericos, newPerifericos...)
	c.IndentedJSON(http.StatusCreated, newPerifericos)
}
