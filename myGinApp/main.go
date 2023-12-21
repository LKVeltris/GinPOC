package main

import (
    "myGinApp/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    routes.InitializeRoutes(router) // Initialize routes

    router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
