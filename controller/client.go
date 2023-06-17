package controller

import (
	"net/http"
	"posts/config"
	"posts/models"
	"posts/repository"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := repository.CreateClient(&client)
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

func GetClientsByFilter(c *gin.Context) {
	var clients []models.Client

	phoneNumberFilter := c.Query("phonenumber")
	ageFilter := c.Query("age")
	tinFilter := c.Query("email")
	fullNameFilter := c.Query("fullname")

	clients, err := repository.GetClients(phoneNumberFilter, ageFilter, tinFilter, fullNameFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &clients)
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

	if err := repository.UpdateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"reason": "Information successfully updated",
	})
}

func DeleteClient(c *gin.Context) {
	var client models.Client
	if err := config.DB.Where("id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "User not found",
		})
		return
	}

	if err := repository.DeleteClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"reason": "Client successfully deleted",
	})
}
