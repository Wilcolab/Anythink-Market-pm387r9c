package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Item struct{
	ID int  `json:"id"`
	NAME string  `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/items", getItems) //Added this line

	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}

func getItems(c *gin.Context){
	items := []Item{
		{ID: 1, NAME: "Galactic Goggles"},
		{ID: 2, Name: "Meteor Muffins"},
		{ID: 3, Name: "Alien Antenna Kit"},
        {ID: 4, Name: "Starlight Lantern"},
        {ID: 5, Name: "Quantum Quill"},
	}

	c.JSON(http.StatusOK, items)
	
}

