package controller

import (
	"test_go/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Ldb *gorm.DB

func LanmarkController(router *gin.Engine, db *gorm.DB) {
	Ldb = db
	router.GET("/landmark", getLandmark)

}

func getLandmark(c *gin.Context) {
	landmarks := []model.Landmark{}
	Ldb.Find(&landmarks)
	c.JSON(200, landmarks)
}
