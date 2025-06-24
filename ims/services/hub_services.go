package services

import (
	
	"context"
	"github.com/angel-omniful/ims/models"

)

func CreateHub(ctx context.Context, hub *models.Hub) error {
	if err := db.Create(hub).Error; err != nil {
		return err
	}
	return nil
}

func GetAllHubs(ctx context.Context) ([]models.Hub, error) {
	var hubs []models.Hub
	if err := db.Find(&hubs).Error; err != nil {
		return nil, err
	}
	return hubs, nil
}

func GetHubByID(ctx context.Context, id string) (models.Hub, error) {
	
	var hub models.Hub
	err := db.First(&hub, id).Error 
	
	return hub,err
}

func UpdateHub(ctx context.Context, id string, updated *models.Hub) error {
	if err := db.Model(&models.Hub{}).Where("id = ?", id).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHub(ctx context.Context, id string) error {
	if err := db.Delete(&models.Hub{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
