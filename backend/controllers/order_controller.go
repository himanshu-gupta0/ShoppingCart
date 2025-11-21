package controllers

import (
	"net/http"

	"backend/database"
	"backend/middleware"
	"backend/models"
	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	CartID uint `json:"cart_id" binding:"required"`
}

// POST /orders
// Converts a cart into an order for the current user.
func CreateOrder(c *gin.Context) {
	user, ok := middleware.GetCurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("id = ? AND user_id = ?", input.CartID, user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart_id for this user"})
		return
	}

	// Create order
	order := models.Order{
		CartID: cart.ID,
		UserID: user.ID,
	}
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create order"})
		return
	}

	// Mark cart as ordered
	cart.Status = "ordered"
	_ = database.DB.Save(&cart).Error

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

// GET /orders
// Lists all orders for the current user.
func ListOrders(c *gin.Context) {
	user, ok := middleware.GetCurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var orders []models.Order
	if err := database.DB.Where("user_id = ?", user.ID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
