package controllers

import (
	"fmt"
	"net/http"

	"backend/database"
	"backend/middleware"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddItemToCartInput struct {
	ItemID uint `json:"item_id" binding:"required"`
}

// POST /carts
// Adds an item to the current user's open cart, creating the cart if needed.
func AddItemToCart(c *gin.Context) {
	user, ok := middleware.GetCurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input AddItemToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure item exists
	var item models.Item
	if err := database.DB.First(&item, input.ItemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	// Find or create open cart for user
	var cart models.Cart
	err := database.DB.Where("user_id = ? AND status = ?", user.ID, "open").First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cart = models.Cart{
				UserID: user.ID,
				Name:   fmt.Sprintf("%s-cart", user.Username),
				Status: "open",
			}
			if err := database.DB.Create(&cart).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create cart"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not find cart"})
			return
		}
	}

	// Create cart item
	cartItem := models.CartItem{
		CartID: cart.ID,
		ItemID: input.ItemID,
	}
	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add item to cart"})
		return
	}

	// Update user's CartID for convenience (not strictly required)
	if user.CartID == nil || *user.CartID != cart.ID {
		user.CartID = &cart.ID
		_ = database.DB.Save(&user).Error
	}

	// Return updated cart with items
	if err := database.DB.Preload("CartItems").First(&cart, cart.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load updated cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"cart": cart})
}

// GET /carts
// Lists all carts for the current user, including cart items.
func ListCarts(c *gin.Context) {
	user, ok := middleware.GetCurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var carts []models.Cart
	if err := database.DB.Preload("CartItems").
		Where("user_id = ?", user.ID).
		Find(&carts).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list carts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"carts": carts})
}
