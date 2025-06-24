package models

import (
	"time"
)

// Inventory represents the inventory table in Postgres.
type Inventory struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID  string    `gorm:"not null;index:idx_inventory_lookup,unique" json:"tenant_id"`
	SellerID  string    `gorm:"not null;index:idx_inventory_lookup,unique" json:"seller_id"`
	HubID     string    `gorm:"type:uuid;not null;index:idx_inventory_lookup,unique" json:"hub_id"`
	SkuID     string    `gorm:"type:uuid;not null;index:idx_inventory_lookup,unique" json:"sku_id"`
	Quantity  int       `gorm:"default:0" json:"quantity"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
func (Inventory) TableName() string {
	return "inventory"
}
