package main

import (
    "net/http"
    "sync"
    "github.com/gin-gonic/gin"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var (
    items []Item
    idCounter int
    mutex sync.Mutex
)

func main() {
    router := gin.Default()
    router.GET("/", greet)
    router.HEAD("/healthcheck", healthcheck)
    router.GET("/items", getItems) // Added ths line 1
    router.POST("/items", addItem) // Added this line 2

    items = []Item{
        {ID: 1, Name: "Galactic Goggles"},
        {ID: 2, Name: "Meteor Muffins"},
        {ID: 3, Name: "Alien Antenna Kit"},
        {ID: 4, Name: "Starlight Lantern"},
        {ID: 5, Name: "Quantum Quill"},
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
    idCounter++
    items = append(items, newItem)

    c.JSON(http.StatusOK, newItem)
}
