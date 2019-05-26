package v1

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/djackreuter/go-bully-api/models"
)

func GetClasses(c *gin.Context) {
    res := models.GetClasses()
    c.JSON(http.StatusOK, res)
}
