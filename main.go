package main

import (
	"infilon_practicaltask/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the router
	router := gin.Default()

	// Defining routes
	router.GET("/person/:personID/info", controller.PersonDetails)
	router.POST("/person/create", controller.CreatePerson)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
