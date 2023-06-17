package repository

import (
	"posts/config"
	"posts/models"
)

func CreateClient(c *models.Client) (uint, error) {
	if err := config.DB.Create(&c).Error; err != nil {
		return 0, err
	}

	return c.ID, nil
}

func GetClients(phoneNumberFilter, ageFilter, tinFilter, fullNameFilter string) ([]models.Client, error) {
	var clients []models.Client
	db := config.DB

	if phoneNumberFilter != "" {
		db = db.Where("phone_number = ?", phoneNumberFilter)
	}
	if ageFilter != "" {
		db = db.Where("age = ?", ageFilter)
	}
	if tinFilter != "" {
		db = db.Where("tin = ?", tinFilter)
	}
	if fullNameFilter != "" {
		db = db.Where("full_name = ?", fullNameFilter)
	}

	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func UpdateClient(c *models.Client) error {
	if err := config.DB.Save(&c).Error; err != nil {
		return err
	}

	return nil
}

func DeleteClient(c *models.Client) error {
	if err := config.DB.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}
