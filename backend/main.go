package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type Items struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.GET("/items", items)
	router.HEAD("/healthcheck", healthcheck)

	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

func items(c *gin.Context) {
	list := []Items{
		{ID: 1, Name: "Galactic Goggles"},
		{ID: 2, Name: "Meteor Muffins"},
		{ID: 3, Name: "Alien Antenna Kit"},
		{ID: 4, Name: "Starlight Lantern"},
		{ID: 5, Name: "Quantum Quill"},
	}

	c.IndentedJSON(http.StatusOK, list)

}


func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
