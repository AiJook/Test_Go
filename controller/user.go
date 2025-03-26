package controller

import (
	"net/http"
	"test_go/dto"
	"test_go/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Customerdb *gorm.DB

func CustomerController(router *gin.Engine, db *gorm.DB) {
	Customerdb = db
	auth := router.Group("/auth")
	{
		auth.POST("/login", login)
	}
	router.PUT("/changePassword", changePassword)
}

func login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := model.Customer{}
	if err := Customerdb.Where("email = ?", input.Email).First(&customer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}

	user := dto.CustomerDTO{}
	copier.Copy(&user, &customer)

	c.JSON(http.StatusOK, user)
}

func changePassword(c *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := model.Customer{}
	if err := Customerdb.Where("email = ?", input.Email).First(&customer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password incorrect"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	customer.Password = string(hashedPassword)
	if err := Customerdb.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}
