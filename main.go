package main

import (
    "github.com/gin-gonic/gin"
    "github.com/djackreuter/go-bully-api/httpd/handler"
)


func main() {
    r := gin.Default()

    r.GET("/", handler.PingGet())

    r.Run()
}
