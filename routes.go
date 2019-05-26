package main

import (
    "github.com/gin-gonic/gin"
    "github.com/djackreuter/go-bully-api/api/v1"
)


func setupRoutes() {
    router := gin.Default()

    r := router.Group("/api/v1")
    {
        r.GET("/class", v1.GetClasses)
    }

    router.Run()
}
