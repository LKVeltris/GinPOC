package routes

import (
    "myGinApp/controllers"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    // Define your routes here
    router.GET("/", controllers.Home)
    router.GET("/users", controllers.GetUsers)
    // Add more routes as needed
}
