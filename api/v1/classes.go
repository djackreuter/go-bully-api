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
        c.JSON(http.StatusNotFound, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func CreateClass(c *gin.Context) {
    var class  models.Class
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    res, err := models.CreateClass(&class)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func UpdateClass(c *gin.Context) {
    var class models.Class
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    res, err := models.UpdateClass(&class)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func DeleteClass(c *gin.Context) {
    class := models.Class{}
    if err := c.ShouldBindJSON(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    if err := models.DeleteClass(&class); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
