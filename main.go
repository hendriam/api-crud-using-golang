package main

import (
	"fmt"
	"shoes/controller"
	"shoes/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	config := lib.LoadConfig()
	logging := lib.LoadLogging()
	url := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	db, err := lib.LoadDatabase()
	if err != nil {
		panic(err)
	}

	ctrl := controller.New(db)

	logging.Info().Msgf("[CONFIG] %v", config)
	logging.Info().Msgf("[APP] Started at => http://%s/", url)

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.GET("/shoes", ctrl.ReadShoes)
	route.GET("/shoes/:shoesId", ctrl.ReadShoesById)
	route.POST("/shoes/create", ctrl.CreateShoes)
	route.PUT("/shoes/update/:shoesId", ctrl.UpdateShoes)
	route.DELETE("/shoes/delete/:shoesId", ctrl.DeleteShoes)

	route.GET("/brands", ctrl.ReadBrands)
	route.GET("/brand/:brandId", ctrl.ReadBrandById)
	route.POST("/brand/create", ctrl.CreateBrand)
	route.PUT("/brand/update/:brandId", ctrl.UpdateBrand)
	route.DELETE("/brand/delete/:brandId", ctrl.DeleteBrand)

	route.Run(url)
}
