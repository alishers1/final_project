package controller

import (
	"net/http"
	"posts/config"
	"posts/models"

	"github.com/gin-gonic/gin"
)

func GetClientsByFilter(c *gin.Context) {
	var clients []models.Client

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

	if err := query.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &clients)
}

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, &client.ID)
}

func UpdateClient(c *gin.Context) {
	var client models.Client
	if err := config.DB.Where("id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Save(&client)

	c.JSON(200, gin.H{
		"reason": "Information successfully updated",
	})
}

func DeleteClient(c *gin.Context) {
	var client models.Client
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"reason": "Client successfully deleted",
	})
}
