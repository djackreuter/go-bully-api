package v1

import (
	"net/http"

	"github.com/djackreuter/go-bully-api/models"
	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {
	res, err := models.GetLocations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetLocation(c *gin.Context) {
	res, err := models.GetLocation(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func AddLocation(c *gin.Context) {
	var newLoc = models.Location{}
	if err := c.ShouldBindJSON(&newLoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
    res, err := models.AddLocation(&newLoc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateLocation(c *gin.Context) {
    loc := models.Location{}
    if err := c.ShouldBindJSON(&loc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    res, err := models.UpdateLocation(&loc)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func DeleteLocation(c *gin.Context) {
    loc := models.Location{}
    if err := c.ShouldBindJSON(&loc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    if err := models.DeleteLocation(&loc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
