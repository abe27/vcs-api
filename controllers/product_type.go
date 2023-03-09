package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/gofiber/fiber/v2"
)

func ProductTypeGetAll(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var prodType models.ProductType
		if err := configs.StoreFormula.Where("FCSKID=?", c.Query("id")).Find(&prodType).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("ProductType with id: %s", c.Query("id"))
		r.Data = &prodType
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var productType []models.ProductType
	if err := configs.StoreFormula.Find(&productType).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All ProductTypes"
	r.Data = &productType
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductTypePost(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ProductType
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var prodType models.ProductType
	prodType.FCCODE = frm.FCCODE
	prodType.FCNAME = frm.FCNAME
	prodType.FCNAME2 = frm.FCNAME2
	if err := configs.StoreFormula.Create(&prodType).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "ProductType Created Successfully"
	r.Data = &prodType
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ProductTypePut(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ProductType
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var prodType models.ProductType
	if err := configs.StoreFormula.Where("FCSKID=?", c.Query("id")).Find(&prodType).Error; err != nil {
		r.Message = fmt.Sprintf("Failed to find product type %s", c.Query("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	prodType.FCCODE = frm.FCCODE
	prodType.FCNAME = frm.FCNAME
	prodType.FCNAME2 = frm.FCNAME2
	if err := configs.StoreFormula.Save(&prodType).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "ProductType Created Successfully"
	r.Data = &prodType
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductTypeDelete(c *fiber.Ctx) error {
	var r models.Response
	var prodType models.ProductType
	if err := configs.StoreFormula.Where("FCSKID=?", c.Query("id")).Find(&prodType).Error; err != nil {
		r.Message = fmt.Sprintf("Failed to find product type %s", c.Query("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.StoreFormula.Delete(&prodType).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Delete ProductType Successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
