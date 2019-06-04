package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/djackreuter/go-bully-api/models"
    "net/http"
)

func GetLocations(c *gin.Context) {
    res, err := models.GetLocations()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}
