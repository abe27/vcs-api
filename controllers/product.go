package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func ProductGetAll(c *fiber.Ctx) error {
	db := configs.StoreFormula
	var r models.Response
	if c.Query("id") != "" {
		var prod models.Product
		if err := db.Preload("ProductType").First(&prod, &models.Product{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Products By %s", c.Query("id"))
		r.Data = &prod
		return c.Status(fiber.StatusOK).JSON(&r)
	}
	var prod []models.Product
	if c.Query("partno") != "" {
		if err := db.Scopes(services.Paginate(c)).
			Preload("ProductType").
			Where("FCCODE LIKE ?", "%"+c.Query("partno")+"%").
			Find(&prod).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Products Like %s", c.Query("id"))
		r.Data = &prod
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Scopes(services.Paginate(c)).Preload("ProductType").Order("FCSKID").Find(&prod).Error; err != nil {
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
	if err := configs.StoreVCST.First(&checkProd, &models.Product{FCCODE: frm.FCCODE}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if checkProd.FCSKID != "" {
		r.Message = fmt.Sprintf("%s Is Duplicate!", frm.FCCODE)
		r.Data = &checkProd
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := configs.StoreVCST.Create(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Create Product %s", frm.FCSKID)
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
	if err := configs.StoreVCST.Where("FCSKID=?", c.Params("id")).First(&prod).Error; err != nil {
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

	if err := configs.StoreVCST.Save(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Update Product %s Completed.", frm.FCSKID)
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductDelete(c *fiber.Ctx) error {
	var r models.Response
	var prod models.Product
	if err := configs.StoreVCST.Where("FCSKID=?", c.Params("id")).First(&prod).Error; err != nil {
		r.Message = fmt.Sprintf("Not Found Product %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.StoreVCST.Delete(&prod).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete Product %s Completed.", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
