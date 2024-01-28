// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route with the separate handler function
	router.GET("/hello", helloHandler)

	// Define a route with a parameter in the URL path
	router.GET("/greet/:name", greetHandler)

	// Define a route for a simple POST API
	router.POST("/post", postHandler)

	// Run the server on port 8080
	router.Run(":8080")
}

// helloHandler is the separate handler function for the "/hello" route
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gin!",
	})
}

// greetHandler is the handler function for the "/greet/:name" route
func greetHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}

// postHandler is the handler function for the "/post" route (POST API)
func postHandler(c *gin.Context) {
	// Parse the JSON payload from the request body
	var requestBody map[string]string
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Access the data from the JSON payload
	message, exists := requestBody["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'message' in the JSON payload",
		})
		return
	}

	// Respond with a message incorporating the data from the JSON payload
	c.JSON(http.StatusOK, gin.H{
		"response": "Received message: " + message,
	})
}
