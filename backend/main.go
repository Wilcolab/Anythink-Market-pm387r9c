package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Views int    `json:"views"`
}

var (
	items = []Item{
		{ID: 1, Name: "Galactic Goggles", Views: 0},
		{ID: 2, Name: "Meteor Muffins", Views: 0},
		{ID: 3, Name: "Alien Antenna Kit", Views: 0},
		{ID: 4, Name: "Starlight Lantern", Views: 0},
		{ID: 5, Name: "Quantum Quill", Views: 0},
	}
	itemsLock sync.Mutex
)

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.GET("/items/:id", getItem)
	router.GET("/items/popular", getPopularItem)

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

func getItems(c *gin.Context) {
	itemsLock.Lock()
	defer itemsLock.Unlock()
	c.JSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	itemsLock.Lock()
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	itemsLock.Unlock()
	c.JSON(http.StatusOK, newItem)
}

func getItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid item ID"})
		return
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()
	for i, item := range items {
		if item.ID == id {
			items[i].Views++
			c.JSON(http.StatusOK, items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func getPopularItem(c *gin.Context) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items found"})
		return
	}

	mostViewed := items[0]
	for _, item := range items {
		if item.Views > mostViewed.Views {
			mostViewed = item
		}
	}
	c.JSON(http.StatusOK, mostViewed)
}
