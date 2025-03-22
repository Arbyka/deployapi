package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Arbyka/deployapi/controller"
)

func main() {
    r := gin.Default()

    r.GET("/api/users", controller.GetUsers)
    r.POST("/api/users", controller.AddUser)

    r.Run() // Default di port 8080
}
