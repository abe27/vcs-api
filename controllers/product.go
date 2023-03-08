package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func ProductGetAll(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	if c.Query("id") != "" {
		var prod models.Product
		if err := db.First(&prod, &models.Product{PRODUCTID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Products By %s", c.Query("id"))
		r.Data = &prod
		return c.Status(fiber.StatusOK).JSON(&r)
	}
	var prod []models.Product
	if err := db.Scopes(services.Paginate(c)).Order("FCSKID").Find(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All Products " + c.Query("id")
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductPost(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Product
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var prod models.Product
	prod.FCTYPE = frm.FCTYPE
	prod.FCCODE = frm.FCCODE
	prod.FCSNAME = frm.FCSNAME
	prod.FCSNAME2 = frm.FCSNAME2
	prod.FCNAME = frm.FCNAME
	prod.FCNAME2 = frm.FCNAME2
	prod.FNAVGCOST = frm.FNAVGCOST
	prod.FNSTDCOST = frm.FNSTDCOST

	var checkProd models.Product
	if err := configs.Store.First(&checkProd, &models.Product{FCCODE: frm.FCCODE}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if checkProd.PRODUCTID != "" {
		r.Message = fmt.Sprintf("%s Is Duplicate!", frm.FCCODE)
		r.Data = &checkProd
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := configs.Store.Create(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Create Product %s", frm.PRODUCTID)
	r.Data = &prod
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ProductPut(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Product
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var prod models.Product
	if err := configs.Store.Where("FCSKID=?", c.Params("id")).First(&prod).Error; err != nil {
		r.Message = fmt.Sprintf("Not Found Product %s", frm.FCCODE)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	prod.FCTYPE = frm.FCTYPE
	prod.FCCODE = frm.FCCODE
	prod.FCSNAME = frm.FCSNAME
	prod.FCSNAME2 = frm.FCSNAME2
	prod.FCNAME = frm.FCNAME
	prod.FCNAME2 = frm.FCNAME2
	prod.FNAVGCOST = frm.FNAVGCOST
	prod.FNSTDCOST = frm.FNSTDCOST

	if err := configs.Store.Save(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Update Product %s Completed.", frm.PRODUCTID)
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductDelete(c *fiber.Ctx) error {
	var r models.Response
	var prod models.Product
	if err := configs.Store.Where("FCSKID=?", c.Params("id")).First(&prod).Error; err != nil {
		r.Message = fmt.Sprintf("Not Found Product %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete Product %s Completed.", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
