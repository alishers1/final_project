package repository

import (
	"posts/config"
	"posts/models"
)

func CreateUser(u *models.User) (uint, error) {
	if err := config.DB.Where(&u).Error; err != nil {
		return 0, err
	}

	return u.ID, nil
}

func GetUsers(ageFilter, emailFilter, fullNameFilter string) ([]models.User, error) {
	var users []models.User

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
		return nil, err
	}

	return users, nil
}

func UpdateUser(u *models.User) error {
	if err := config.DB.Save(&u).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(u *models.User) error {
	if err := config.DB.Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
