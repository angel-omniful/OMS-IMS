package models

import "time"


type Hub struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID  string    `gorm:"not null" json:"tenant_id"`
	SellerID  string    `gorm:"not null" json:"seller_id"`
	Name      string    `gorm:"not null" json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
