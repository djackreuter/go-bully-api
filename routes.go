package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/djackreuter/go-bully-api/api/v1"
)


func setupRoutes() {
    r := gin.Default()

    g1 := r.Group("/api/v1")
    {
        g1.GET("/classes", v1.GetClasses)
        g1.GET("/class/:id", v1.GetClass)
        g1.POST("/classes", v1.CreateClass)
    }

    r.Run()
    fmt.Println("Listening on port 8080")
}
