package controller

import (
	"infilon_practicaltask/models"
	"infilon_practicaltask/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var newPerson models.PersonRequest

	// Bind the JSON input to the newPerson struct
	if err := c.BindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Creating a new person using the service
	err := service.CreatePerson(newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}
