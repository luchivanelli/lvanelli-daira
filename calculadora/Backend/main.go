package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type operation struct {
	Operation_ID int    `json:"id"`
	Operation    string `json:"operation"`
	Result       string `json:"result"`
}

var operations = []operation{
	{Operation_ID: 1, Operation: "10-3", Result: "7"},
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func postData(c *gin.Context) {
	var newOperation operation
	if err := c.BindJSON(&newOperation); err != nil {
		return
	}

	operations = append(operations, newOperation)

	c.IndentedJSON(http.StatusCreated, operations)
}


// func getOperationID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range operations {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Operation not found"})
// }

func main() {
	fmt.Println(operations)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/operations", getData)
	router.POST("/operations", postData)
	// router.GET("/operations/:id", getOperationID)

	router.Run("localhost:5173")
}
