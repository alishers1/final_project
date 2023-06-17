package controller

import (
	"net/http"
	"posts/config"
	"posts/models"
	"posts/repository"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "Couldn't read from body",
		})
		return
	}

	id, err := repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"id": id,
	})
}

func GetUsersByFilter(c *gin.Context) {
	var users []models.User

	ageFilter := c.Query("age")
	emailFilter := c.Query("email")
	fullNameFilter := c.Query("fullname")

	
	users, err := repository.GetUsers(ageFilter, emailFilter, fullNameFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "Error while reading the body",
		})
		return
	}

	if err := repository.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &user.ID)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "User not found",
		})
		return
	}

	if err := repository.DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reason": "User successfully deleted",
	})
}
