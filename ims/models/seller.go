package models

import (
	"time"
)
type Seller struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID  string    `gorm:"type:uuid;not null" json:"tenant_id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
