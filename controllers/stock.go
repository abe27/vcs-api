package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func StockGetAll(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var stock models.Stock
		if err := configs.StoreFormula.
			Preload("Whs").
			Preload("Product.ProductType").
			First(&stock, &models.Stock{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Stock By %s", c.Query("id"))
		r.Data = &stock
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var stock []models.Stock
	if err := configs.StoreFormula.
		Scopes(services.Paginate(c)).
		Preload("Whs").
		Preload("Product.ProductType").
		Find(&stock).
		Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All Stock"
	r.Data = &stock
	return c.Status(fiber.StatusOK).JSON(&r)
}

func StockPost(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func StockPut(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func StockDelete(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
