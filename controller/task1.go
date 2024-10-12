package controller

import (
	"infilon_practicaltask/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PersonDetails(c *gin.Context) {
	// personID from the URL parameters
	personIDStr := c.Param("personID")
	if personIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person ID cannot be empty"})
		return
	}

	// Convert personID to integer
	personID, err := strconv.Atoi(personIDStr)
	if err != nil || personID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	// Fetch person details from the service
	details, err := service.GetPersonDetails(personID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if details == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, details)
}
