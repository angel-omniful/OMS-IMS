package services

import (
	"context"
	"time"
	"log"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/myDb"
	"github.com/omniful/go_commons/redis"
	"errors"
	"gorm.io/gorm"
)

var cache *redis.Client=myDb.GetCache()


func CreateInventory(ctx context.Context, inv *models.Inventory) error {
	err:=db.Create(inv).Error
	if err != nil {
		return err
	}
    key:=inv.SkuID + ":" + inv.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error setting cache for inventory:", err1)
	}
	return nil
}
func GetAllInventory(ctx context.Context) ([]models.Inventory, error) {
	var inventory []models.Inventory
	err := db.Find(&inventory).Error
	return inventory, err
}

func GetInventoryByID(ctx context.Context, id string) (models.Inventory, error) {
	var inv models.Inventory
	
	err := db.First(&inv, "id = ?", id).Error

	key:=inv.SkuID + ":" + inv.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error setting cache for inventory:", err1)
	}
	return inv, err
}

func UpdateInventory(ctx context.Context, id string, updated *models.Inventory) error {
	err:=db.Model(&models.Inventory{}).Where("id = ?", id).Updates(updated).Error
	if err != nil {
		return err	
	}
	key := updated.SkuID + ":" + updated.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error updating cache for inventory:", err1)
	}	
	return nil	
}

func DeleteInventory(ctx context.Context, id string) error {
	inv,err:=GetInventoryByID(ctx, id)

	if err != nil {
		log.Println("Error fetching inventory for deletion:", err)
		return err
	}

	key:=inv.SkuID + ":" + inv.HubID

	_,err1:=cache.Del(ctx, key)
	if err1 != nil {
		log.Println("Error deleting cache for inventory:", err1)
	}

	err2:=db.Delete(&models.Inventory{}, "id = ?", id).Error
	return err2
}

func ValidateInventory(ctx context.Context, skuID string, hubID string) (bool, error) {
	key := skuID + ":" + hubID

	_, err := cache.Get(ctx, key)
	if err != nil {
		var order models.Inventory
		dbErr := db.Where("hub_id = ? AND sku_id = ?", hubID, skuID).First(&order).Error
		if dbErr != nil {
			// Return false if record not found or some other DB error
			if errors.Is(dbErr, gorm.ErrRecordNotFound) {
				return false, nil
			}
			return false, dbErr
		}
	}

	return true, nil
}

func CheckInventory(ctx context.Context, skuID string, hubID string, qty int) (bool, error) {
	var inv models.Inventory

	// Check inventory
	err := db.Where("sku_id = ? AND hub_id = ? AND quantity >= ?", skuID, hubID, qty).First(&inv).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // No inventory found or insufficient
		}
		return false, err // DB or other error
	}

	// Deduct quantity
	inv.Quantity -= qty
	if err := db.Save(&inv).Error; err != nil {
		return false, err
	}

	return true, nil
}
