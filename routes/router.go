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

	// Product Type Router
	prodType := r.Group("/productType")
	prodType.Get("", controllers.ProductTypeGetAll)
	prodType.Post("", controllers.ProductTypePost)
	prodType.Put("/:id", controllers.ProductTypePut)
	prodType.Delete("/:id", controllers.ProductTypeDelete)

	// Product Router
	prod := r.Group("/product")
	prod.Get("", controllers.ProductGetAll)
	prod.Post("", controllers.ProductPost)
	prod.Put("/:id", controllers.ProductPut)
	prod.Delete("/:id", controllers.ProductDelete)

	// Whs Router
	whs := r.Group("/whs")
	whs.Get("", controllers.WhsGetAll)
	whs.Post("", controllers.WhsPost)
	whs.Put("/:id", controllers.WhsPut)
	whs.Delete("/:id", controllers.WhsDelete)

	// Stock Router
	stock := r.Group("/stock")
	stock.Get("", controllers.StockGetAll)
	stock.Post("", controllers.StockPost)
	stock.Put("/:id", controllers.StockPut)
	stock.Delete("/:id", controllers.StockDelete)

	// GlRef Router
	glref := r.Group("/glref")
	glref.Get("", controllers.GlRefGetAll)
	glref.Post("", controllers.GlRefPost)
	glref.Put("/:id", controllers.GlRefPut)
	glref.Delete("/:id", controllers.GlRefDelete)

	// GlRef Router
	refprod := r.Group("/refprod")
	refprod.Get("", controllers.RefProdGetAll)
	refprod.Post("", controllers.RefProdPost)
	refprod.Put("/:id", controllers.RefProdPut)
	refprod.Delete("/:id", controllers.RefProdDelete)

	// Use Router Middleware
	app := r.Use(services.AuthorizationRequired)
	app.Get("/employee/all", controllers.GetAllEmployee)
	app.Get("/employee/me", controllers.GetProfile)
}
