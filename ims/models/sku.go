package models

import (
	"time"
)

type SKU struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID  string    `gorm:"not null;index:idx_tenant_seller_sku,unique" json:"tenant_id"`
	SellerID  string    `gorm:"not null;index:idx_tenant_seller_sku,unique" json:"seller_id"`
	SKUCode   string    `gorm:"not null;index:idx_tenant_seller_sku,unique" json:"sku_code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}


