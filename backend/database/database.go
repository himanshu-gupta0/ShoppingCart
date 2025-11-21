package database

import (
	"log"

	"backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// Init initializes the database connection and performs migrations.
func Init() error {
	var err error
	DB, err = gorm.Open("sqlite3", "shopping_cart.db")
	if err != nil {
		return err
	}

	// Auto-migrate all models
	if err = DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	).Error; err != nil {
		return err
	}

	log.Println("database initialized and migrated")
	return nil
}

// Close closes the DB connection.
func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}
}
