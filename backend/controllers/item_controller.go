package controllers

import (
	"net/http"

	"backend/database"
	"backend/models"
	"github.com/gin-gonic/gin"
)

type CreateItemInput struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"` // e.g. "available"
}

// POST /items
func CreateItem(c *gin.Context) {
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name:   input.Name,
		Status: input.Status,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"item": item})
}

// GET /items
func ListItems(c *gin.Context) {
	var items []models.Item
	if err := database.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
