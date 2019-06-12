package v1

import (
	"net/http"

	"github.com/djackreuter/go-bully-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	res, err := models.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetCliqueStudents(c *gin.Context) {
	res, err := models.GetCliqueStudents(c.Param("clique_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetCliqueStudent(c *gin.Context) {
	res, err := models.GetCliqueStudent(c.Param("clique_id"), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateCliqueStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res, err := models.CreateCliqueStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateCliqueStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res, err := models.UpdateCliqueStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteCliqueStudent(c *gin.Context) {
	err := models.DeleteCliqueStudent(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
