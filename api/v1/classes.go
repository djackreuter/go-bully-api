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

func GetClass(c *gin.Context) {
    res, err := models.GetClass(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, err)
    } else {
        c.JSON(http.StatusOK, res)
    }
}

