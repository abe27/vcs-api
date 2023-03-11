package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func RefProdGetAll(c *fiber.Ctx) error {
	var r models.Response
	var refprod []models.RefProd
	if c.Query("glref") != "" {
		if err := configs.StoreFormula.
			Preload("GlRef.FromWhs").
			Preload("GlRef.ToWhs").
			Preload("Product.ProductType").
			Preload("Whs").
			Find(&refprod, &models.RefProd{FCGLREF: c.Query("glref")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show REFPROD By ID: %s", c.Query("glref"))
		r.Data = &refprod
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("fcprod") != "" {
		if err := configs.StoreFormula.
			Preload("GlRef.FromWhs").
			Preload("GlRef.ToWhs").
			Preload("Product.ProductType").
			Preload("Whs").
			Find(&refprod, &models.RefProd{FCPROD: c.Query("fcprod")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show REFPROD By ID: %s", c.Query("fcprod"))
		r.Data = &refprod
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := configs.StoreFormula.
		Scopes(services.Paginate(c)).
		Preload("GlRef.FromWhs").
		Preload("GlRef.ToWhs").
		Preload("Product.ProductType").
		Preload("Whs").
		Find(&refprod).
		Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All REFPROD"
	r.Data = &refprod
	return c.Status(fiber.StatusOK).JSON(&r)
}

func RefProdPost(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func RefProdPut(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func RefProdDelete(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
