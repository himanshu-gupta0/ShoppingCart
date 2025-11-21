package models

import "time"

type Cart struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"` // e.g. "open", "ordered"
	CreatedAt time.Time  `json:"created_at"`

	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	ID     uint `gorm:"primary_key" json:"id"`
	CartID uint `json:"cart_id"`
	ItemID uint `json:"item_id"`
}
