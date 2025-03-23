package controller

import (
	"net/http"
	"test_go/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Cdb *gorm.DB

func CountryController(router *gin.Engine, db *gorm.DB) {
	Cdb = db
	router.GET("/country", getCountry)
	router.GET("/country/:id", getCountrybyID)
	router.GET("/country/name", getCountrybyName)
	router.POST("/country", insertCountry)
	router.PUT("/country/:id", updatCountrybyID)
	router.PUT("/country", updatCountry)


}

func getCountry(c *gin.Context) {
	country := []model.Country{}
	Cdb.Find(&country)
	c.JSON(200, country)

}
func getCountrybyID(c *gin.Context) {
	id := c.Param("id")
	country := model.Country{}
	if err := Cdb.First(&country, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}
	c.JSON(200, country)
}
func getCountrybyName(c *gin.Context) {
	name := c.Query("name")
	country := []model.Country{}
	if err := Cdb.Where("Name LIKE ?", "%"+name+"%").Find(&country).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}
	c.JSON(200, country)
}

func insertCountry(c *gin.Context) {
	country := model.Country{}
	c.ShouldBindJSON(&country)
	Cdb.Create(&country)
	c.JSON(200, country)
}

func updatCountrybyID(c *gin.Context) {
	country := model.Country{}
	c.ShouldBindJSON(&country)
	Cdb.Save(&country)
	c.JSON(200, country)

}
func updatCountry(c *gin.Context) {
	var countries []model.Country
	c.ShouldBindJSON(&countries)
	for _, country := range countries {
		Cdb.Model(&model.Country{}).Where("idx = ?", country.Idx).Updates(&country)
	}
	c.JSON(200, countries)
}
