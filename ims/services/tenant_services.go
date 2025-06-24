package services

import (
	"context"

	"github.com/angel-omniful/ims/models"

)



func CreateTenant(ctx context.Context, tenant *models.Tenant) error {
	return db.Create(tenant).Error
}

func GetAllTenants(ctx context.Context) ([]models.Tenant, error) {
	var tenants []models.Tenant
	err := db.Find(&tenants).Error
	return tenants, err
}

func GetTenantByID(ctx context.Context, id string) (models.Tenant, error) {
	var tenant models.Tenant
	err := db.First(&tenant, "id = ?", id).Error
	return tenant, err
}

func UpdateTenant(ctx context.Context, id string, updated *models.Tenant) error {
	return db.Model(&models.Tenant{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteTenant(ctx context.Context, id string) error {
	return db.Delete(&models.Tenant{}, "id = ?", id).Error
}
