package controllers

import (
	"fmt"
	"strings"

	"github.com/abe27/cvst20/api/configs"
	"github.com/abe27/cvst20/api/models"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllEmployee(c *fiber.Ctx) error {
	var r models.Response
	var emp []models.EMPLR
	err := configs.Store.Find(&emp).Error
	if err != nil {
		r.Message = err.Error()
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Data = &emp
	return c.Status(fiber.StatusOK).JSON(&r)
}

func EmployeeLogin(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmLogin
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	if len(frm.UserName) <= 0 || len(frm.Password) <= 0 {
		r.Message = "Please enter username and password"
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	db := configs.Store
	var emp models.EMPLR
	err := db.Where("FCLOGIN=?", strings.ToUpper(frm.UserName)).Where("FCPW=?", strings.ToUpper(frm.Password)).First(&emp).Error
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	err = configs.Store.First(&emp).Error
	if err != nil {
		r.Message = err.Error()
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Create JWT token
	auth := services.CreateToken(emp)
	r.Message = "Login Success"
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetProfile(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Get Profile"

	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = "Unauthorized"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}
	empID, er := services.ValidateToken(token)
	if er != nil {
		r.Message = er.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	id := fmt.Sprintf("%v", empID)
	var emp models.EMPLR
	if err := configs.Store.First(&emp, "FCSKID", id).Error; err != nil {
		r.Message = er.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = fmt.Sprintf("Show Profile: %v", id)
	r.Data = &emp
	return c.Status(fiber.StatusOK).JSON(&r)
}
