package services

import (
	"context"

	"github.com/angel-omniful/ims/models"
	
)



func CreateSeller(ctx context.Context, seller *models.Seller) error {
	if err := db.Create(seller).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSellers(ctx context.Context) ([]models.Seller, error) {
	var sellers []models.Seller
	if err := db.Find(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

func GetSellerByID(ctx context.Context, id string) (models.Seller, error) {
	var seller models.Seller
	err := db.First(&seller, "id = ?", id).Error // safer to use named query
	return seller, err
}

func UpdateSeller(ctx context.Context, id string, updated *models.Seller) error {
	if err := db.Model(&models.Seller{}).Where("id = ?", id).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSeller(ctx context.Context, id string) error {
	if err := db.Delete(&models.Seller{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
