package v1

import (
	"net/http"

	"github.com/djackreuter/go-bully-api/models"
	"github.com/gin-gonic/gin"
)

func GetMissions(c *gin.Context) {
    res, err := models.GetMissions()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func GetMission(c *gin.Context) {
    res, err := models.GetMission(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func CreateMission(c *gin.Context) {
    var mission models.Mission
    if err := c.ShouldBindJSON(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    res, err := models.CreateMission(&mission)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func UpdateMission(c *gin.Context) {
    var mission models.Mission
    if err := c.ShouldBindJSON(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err}) 
        return
    }
    res, err := models.UpdateMission(&mission)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, res)
}

func DeleteMission(c *gin.Context) {
    var mission models.Mission
    if err := c.ShouldBindJSON(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    if err := models.DeleteMission(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Mission deleted successfully"})
}
