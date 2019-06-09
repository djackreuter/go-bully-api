package v1

import (
	"net/http"

	"github.com/djackreuter/go-bully-api/models"
	"github.com/gin-gonic/gin"
)

func GetCliques(c *gin.Context) {
    res, err := models.GetCliques()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func GetClique(c *gin.Context) {
    res, err := models.GetClique(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func CreateClique(c *gin.Context) {
    var clique models.Clique
    if err := c.ShouldBindJSON(&clique); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    res, err := models.CreateClique(&clique)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func UpdateClique(c *gin.Context) {
    var clique models.Clique
    if err := c.ShouldBindJSON(&clique); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err}) 
        return
    }
    res, err := models.UpdateClique(&clique)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func DeleteClique(c *gin.Context) {
    var clique models.Clique
    if err := c.ShouldBindJSON(&clique); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    if err := models.DeleteClique(&clique); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Clique deleted successfully"})
}
