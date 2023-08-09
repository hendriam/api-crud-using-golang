package controller

import (
	"net/http"
	"shoes/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app App) ReadShoes(c *gin.Context) {
	Brands, err := app.model.SelectShoes()
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

func (app App) ReadShoesById(c *gin.Context) {
	shoesId, _ := strconv.Atoi(c.Param("shoesId"))

	Shoes, err := app.model.SelectShoesById(shoesId)
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	app.logging.Info().Msgf("[DATA] => %+v\n", Shoes)

	c.JSON(http.StatusOK, gin.H{
		"data":    Shoes,
		"success": true,
	})
}

func (app App) CreateShoes(c *gin.Context) {
	var shoes model.Shoes
	if err := c.ShouldBindJSON(&shoes); err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.InsertShoes(shoes)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	shoes.Id = int(lastId)

	app.logging.Info().Msgf("[CREATE] success => %+v\n", shoes)

	c.JSON(http.StatusOK, gin.H{
		"data":    shoes,
		"success": true,
	})
}

func (app App) UpdateShoes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("shoesId"))
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	var shoes model.Shoes

	if err := c.ShouldBindJSON(&shoes); err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// check if row exist
	shoesIfExist, err := app.model.SelectShoesById(id)
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.UpdateShoes(shoesIfExist.Id, shoes)
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
			"error":   lastId,
			"success": false,
		})
		return
	}

	shoes.Id = id

	app.logging.Info().Msgf("[UPDATE] success => %+v\n", shoes)

	c.JSON(http.StatusOK, gin.H{
		"data":    shoes,
		"success": true,
	})
}

func (app App) DeleteShoes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("shoesId"))
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// chect if row exist
	shoesIfExist, err := app.model.SelectShoesById(id)
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	lastId, err := app.model.DeleteShoes(shoesIfExist.Id)
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

	app.logging.Info().Msgf("[DELETE] success,  %+v\n", shoesIfExist)

	c.JSON(http.StatusOK, gin.H{
		"data":    shoesIfExist,
		"success": true,
	})
}
