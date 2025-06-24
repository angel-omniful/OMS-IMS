package model

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID       string             `json:"_id,omitempty" bson:"_id,omitempty"`
	TenantID string             `json:"tenant_id,omitempty" bson:"tenant_id,omitempty"`
	SellerID string             `json:"seller_id,omitempty" bson:"seller_id,omitempty"`
	HubID    string             `json:"hub_id,omitempty" bson:"hub_id,omitempty"`
	SkuID    string             `json:"sku_id,omitempty" bson:"sku_id,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	Quantity int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

