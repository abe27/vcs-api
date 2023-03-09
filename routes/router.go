package routes

import (
	"github.com/abe27/cvst20/api/controllers"
	"github.com/abe27/cvst20/api/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HandlerHello)

	// Group Prefix Router
	r := c.Group("/api/v1")
	user := r.Group("/employee")
	user.Post("", controllers.EmployeeLogin)

	// Use Router Middleware
	app := r.Use(services.AuthorizationRequired)
	app.Get("/employee/all", controllers.GetAllEmployee)
	app.Get("/employee/me", controllers.GetProfile)

	// Product Type Router
	prodType := app.Group("/productType")
	prodType.Get("", controllers.ProductTypeGetAll)
	prodType.Post("", controllers.ProductTypePost)
	prodType.Put("/:id", controllers.ProductTypePut)
	prodType.Delete("/:id", controllers.ProductTypeDelete)

	// Product Router
	prod := app.Group("/product")
	prod.Get("", controllers.ProductGetAll)
	prod.Post("", controllers.ProductPost)
	prod.Put("/:id", controllers.ProductPut)
	prod.Delete("/:id", controllers.ProductDelete)
}
