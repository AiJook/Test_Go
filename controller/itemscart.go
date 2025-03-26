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
	router.POST("/search", searchItems)
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

func searchItems(c *gin.Context) {
	var input struct {
		Description string  `json:"description"`
		MinPrice    float64 `json:"min_price"`
		MaxPrice    float64 `json:"max_price"`
	}

	// รับค่า JSON ที่ส่งมา
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var items []model.Product
	// สร้าง query ตามเงื่อนไขทั้งหมด
	query := Idb.Model(&items)
	query = query.Where("description LIKE ? AND price BETWEEN ? AND ?", "%"+input.Description+"%", input.MinPrice, input.MaxPrice)

	// ค้นหาผลลัพธ์
	if err := query.Find(&items).Error; err != nil {
		c.JSON(404, gin.H{"error": "No products found"})
		return
	}

	// ส่งผลลัพธ์กลับไป
	c.JSON(200, items)
}
