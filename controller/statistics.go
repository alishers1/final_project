package controller

import (
	"net/http"
	"posts/config"
	"posts/models"

	"github.com/gin-gonic/gin"
)

func ClientStats(c *gin.Context) {
	var totalClients int64
	var maleClients int64
	var femaleClients int64
	var clientsWithProducts int64

	config.DB.Model(&models.Client{}).Count(&totalClients)
	config.DB.Model(&models.Client{}).Where("gender = ?", "Male").Count(&maleClients)
	config.DB.Model(&models.Client{}).Where("gender = ?", "Female").Count(&femaleClients)
	config.DB.Model(&models.Client{}).Joins("JOIN produtcs ON client.id = products.client_id").Count(&clientsWithProducts)

	c.JSON(http.StatusOK, gin.H{
		"totalClients":        totalClients,
		"maleClients":         maleClients,
		"femaleClients":       femaleClients,
		"clientsWithProducts": clientsWithProducts,
	})
}
