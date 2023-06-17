package controller

import (
	"net/http"
	"posts/config"
	"posts/models"

	"github.com/gin-gonic/gin"
)

func GetUsersByFilter(c *gin.Context) {
	var users []models.User

	ageFilter := c.Query("age")
	emailFilter := c.Query("email")
	fullNameFilter := c.Query("fullname")

	query := config.DB
	if ageFilter != "" {
		query = query.Where("age = ?", ageFilter)
	}
	if emailFilter != "" {
		query = query.Where("email = ?", emailFilter)
	}
	if fullNameFilter != "" {
		query = query.Where("full_name LIKE ?", fullNameFilter+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "Couldn't read from body",
		})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Couldn't create the user",
		})
		return
	}

	c.JSON(201, &user.ID)
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

	config.DB.Save(&user)

	c.JSON(http.StatusOK, &user.ID)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reason": "User successfully deleted",
	})
}
