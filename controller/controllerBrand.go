package controller

import (
	"net/http"
	"shoes/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app App) ReadBrands(c *gin.Context) {
	Brands, err := app.model.SelectBrand()
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	app.logging.Info().Msgf("[DATA] => %+v\n", Brands)

	c.JSON(http.StatusOK, gin.H{
		"data":    Brands,
		"success": true,
	})
}

func (app App) ReadBrandById(c *gin.Context) {
	brandId, _ := strconv.Atoi(c.Param("brandId"))

	Brand, err := app.model.SelectBrandById(brandId)
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	app.logging.Info().Msgf("[DATA] => %+v\n", Brand)

	c.JSON(http.StatusOK, gin.H{
		"data":    Brand,
		"success": true,
	})
}

func (app App) CreateBrand(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.InsertBrand(brand)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	brand.Id = int(lastId)

	app.logging.Info().Msgf("[CREATE] success => %+v\n", brand)

	c.JSON(http.StatusOK, gin.H{
		"data":    brand,
		"success": true,
	})
}

func (app App) UpdateBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("brandId"))
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	var brand model.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// chect if row exist
	brandIfExist, err := app.model.SelectBrandById(id)
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.UpdateBrand(brandIfExist.Id, brand)
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	if lastId == 0 {
		app.logging.Error().Msgf("[UPDATE] failed  %+v\n", lastId)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "update failed",
			"success": false,
		})
		return
	}

	brand.Id = id

	app.logging.Info().Msgf("[UPDATE] success => %+v\n", brand)

	c.JSON(http.StatusOK, gin.H{
		"data":    brand,
		"success": true,
	})
}

func (app App) DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("brandId"))
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// chect if row exist
	brandIfExist, err := app.model.SelectBrandById(id)
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.DeleteBrand(brandIfExist.Id)
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	if lastId == 0 {
		app.logging.Error().Msgf("[DELETE] failed  %+v\n", lastId)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "delete failed",
			"success": false,
		})
		return
	}

	app.logging.Info().Msgf("[DELETE] success,  %+v\n", brandIfExist)

	c.JSON(http.StatusOK, gin.H{
		"data":    brandIfExist,
		"success": true,
	})
}
