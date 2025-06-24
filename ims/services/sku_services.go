package services

import (
	"github.com/angel-omniful/ims/myContext"
	"context"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/myDb"
	"gorm.io/gorm"
)

//my context
var ctx = myContext.GetContext()
//master db
var db *gorm.DB= myDb.GetDb().GetMasterDB(ctx)

func CreateSKU(ctx context.Context, sku *models.SKU) error {
	if err := db.Create(sku).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSKUs(ctx context.Context) ([]models.SKU, error) {
	var skus []models.SKU
	if err := db.Find(&skus).Error; err != nil {
		return nil, err
	}
	return skus, nil
}

func GetSKUByID(ctx context.Context, id string) (models.SKU, error) {
	
	var sku models.SKU
	err := db.First(&sku, id).Error // Only if `id` is the primary key
	
	return sku,err
}

func UpdateSKU(ctx context.Context, id string, updated *models.SKU) error {
	if err := db.Model(&models.SKU{}).Where("id = ?", id).Updates(updated).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSKU(ctx context.Context, id string) error {
	if err := db.Delete(&models.SKU{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func GetSKUsByTenantID(ctx context.Context, tenantID string) ([]models.SKU, error) {
	var skus []models.SKU
	err := db.Where("tenant_id = ?", tenantID).Find(&skus).Error
	return skus, err
}

func GetSKUsBySellerID(ctx context.Context, sellerID string) ([]models.SKU, error) {
	var skus []models.SKU
	err := db.Where("seller_id = ?", sellerID).Find(&skus).Error
	return skus, err
}

func GetSKUsBySKUCode(ctx context.Context, skuCode string) ([]models.SKU, error) {
	var skus []models.SKU
	err := db.Where("sku_code = ?", skuCode).Find(&skus).Error
	return skus, err
}
