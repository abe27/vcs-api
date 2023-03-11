package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func GlRefGetAll(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var glref models.GlRef
		if err := configs.StoreFormula.
			Preload("FromWhs").
			Preload("ToWhs").
			First(&glref, &models.GlRef{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show GlRef By ID: %s", glref.ID)
		r.Data = &glref
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var glref []models.GlRef
	if err := configs.StoreFormula.
		Scopes(services.Paginate(c)).
		Preload("FromWhs").
		Preload("ToWhs").
		Find(&glref).
		Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All GLREF"
	r.Data = &glref
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlRefPost(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlRefPut(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlRefDelete(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
