package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dsn := viper.GetString("mysql.dsn")
	dialetor := mysql.Open(dsn)

	db, err := gorm.Open(dialetor)
	if err != nil {
		panic(err)
	}

	fmt.Printf("connection successfully")
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	LanmarkController(router, db)
	CountryController(router, db)

	router.Run()

}
