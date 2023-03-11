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
			// Preload("RefProd.GlRef.FromWhs").
			// Preload("RefProd.GlRef.ToWhs").
			First(&stock, &models.Stock{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Stock By %s", c.Query("id"))
		r.Data = &stock
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var stock []models.Stock
	if c.Query("partno") != "" {
		var part []models.Product
		if err := configs.StoreFormula.Scopes(services.Paginate(c)).Select("FCSKID").Where("FCCODE like ?", "%"+c.Query("partno")+"%").Find(&part).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		var lstPart []string
		for _, p := range part {
			lstPart = append(lstPart, p.FCSKID)
		}

		if err := configs.StoreFormula.
			Scopes(services.Paginate(c)).
			Preload("Whs").
			Preload("Product.ProductType").
			// Preload("RefProd.GlRef.FromWhs").
			// Preload("RefProd.GlRef.ToWhs").
			Where("FCPROD in ?", lstPart).
			Find(&stock).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All Stock"
		r.Data = &stock
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := configs.StoreFormula.
		Scopes(services.Paginate(c)).
		Preload("Whs").
		Preload("Product.ProductType").
		// Preload("RefProd.GlRef.FromWhs").
		// Preload("RefProd.GlRef.ToWhs").
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
