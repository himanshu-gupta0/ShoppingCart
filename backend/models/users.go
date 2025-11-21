package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Password  string     `gorm:"not null" json:"-"` // not exposed via JSON
	Token     *string    `json:"-"`                 // active login token
	CartID    *uint      `json:"cart_id"`
	CreatedAt time.Time  `json:"created_at"`
}
