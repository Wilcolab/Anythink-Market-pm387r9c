package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ViewCount int    `json:"viewCount"`
}

var (
	items     []Item
	idCounter int
	mutex     sync.Mutex
)

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.GET("/items/:id", getItemByID) 

	items = []Item{
		{ID: 1, Name: "Galactic Goggles", ViewCount: 0},
		{ID: 2, Name: "Meteor Muffins", ViewCount: 0},
		{ID: 3, Name: "Alien Antenna Kit", ViewCount: 0},
		{ID: 4, Name: "Starlight Lantern", ViewCount: 0},
		{ID: 5, Name: "Quantum Quill", ViewCount: 0},
	}
	idCounter = 6

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
	c.JSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	newItem.ID = idCounter
	newItem.ViewCount = 0
	idCounter++
	items = append(items, newItem)

	c.JSON(http.StatusOK, newItem)
}

func getItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item *Item

	for i := range items {
		if items[i].ID == id {
			item = &items[i]
			break
		}
	}

	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	go incrementViewCount(item)

	c.JSON(http.StatusOK, item)
}

func incrementViewCount(item *Item) {
	mutex.Lock()
	defer mutex.Unlock()

	item.ViewCount++
	time.Sleep(100 * time.Millisecond) 
}
