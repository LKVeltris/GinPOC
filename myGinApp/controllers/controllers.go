package controllers

import (
    "net/http"
    "myGinApp/models"
    "github.com/gin-gonic/gin"
)

// Home is a simple controller to test the setup
func Home(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Welcome to the Gin web app!"})
}

// GetUsers fetches all users from the database
func GetUsers(c *gin.Context) {
    users, err := models.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": users})
}