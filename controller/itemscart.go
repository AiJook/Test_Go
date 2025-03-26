package controller

import (
	"test_go/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Idb *gorm.DB

func ItemCartController(router *gin.Engine, db *gorm.DB) {
	Idb = db
	router.GET("/itemscart/:id", itemsCart)
}

func itemsCart(c *gin.Context) {
	userID := c.Param("id")

	var carts []model.Cart
	if err := Idb.Preload("CartItems").Preload("CartItems.Product").Where("customer_id = ?", userID).Find(&carts).Error; err != nil {
		c.JSON(404, gin.H{"error": "No items found for the given user ID"})
		return
	}

	c.JSON(200, carts)
}
