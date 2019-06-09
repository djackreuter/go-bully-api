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
        g1.PATCH("/classes", v1.UpdateClass)
        g1.DELETE("/class", v1.DeleteClass)

        g1.GET("/locations", v1.GetLocations)
        g1.GET("/location/:id", v1.GetLocation)
        g1.POST("/location", v1.AddLocation)
        g1.PATCH("/location", v1.UpdateLocation)
        g1.DELETE("/location", v1.DeleteLocation)
    }

    r.Run()
    fmt.Println("Listening on port 8080")
}
