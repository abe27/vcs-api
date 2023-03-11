package controllers

import (
	"fmt"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/gofiber/fiber/v2"
)

func WhsGetAll(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var whs models.Whs
		if err := configs.StoreFormula.Where("FCSKID=?", c.Query("id")).Find(&whs).Error; err != nil {
			r.Message = fmt.Sprintf("%s is not found", c.Query("id"))
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = fmt.Sprintf("Show Whs ID: %s", c.Query("id"))
		r.Data = &whs
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var whs []models.Whs
	if err := configs.StoreFormula.Order("FCCODE").Find(&whs).Error; err != nil {
		r.Message = fmt.Sprintf("%s is not found", c.Query("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = fmt.Sprintf("Show Whs ID: %s", c.Query("id"))
	r.Data = &whs
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsPost(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsPut(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsDelete(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
