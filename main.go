package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pat struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Symp string `json:"symp"`
}

var temp int = 0
var cur int = 2
var queue = []pat{
	{Id: 1, Name: "ram", Age: 45, Symp: "asthama"},
	{Id: 2, Name: "rohit", Age: 62, Symp: "cancer"},
	{Id: 3, Name: "aditya", Age: 56, Symp: "fever"},
}

func dequeue() int {
	if temp > cur {
		return -1
	}
	x := temp
	temp++
	return x
}

func nextpat(c *gin.Context) {
	x := dequeue()
	if x == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Patient left"})
		return
	}
	c.IndentedJSON(http.StatusOK, queue[x])
}

func newpatient(c *gin.Context) {
	var newpat pat

	if err := c.Bind(&newpat); err != nil {
		return
	}
	cur++
	newpat.Id = cur + 1
	queue = append(queue, newpat)
	c.IndentedJSON(http.StatusCreated, newpat.Id)
}
func main() {
	router := gin.Default()
	router.GET("/pat", nextpat)
	router.POST("/pat", newpatient)
	router.Run("localhost:8080")
}
