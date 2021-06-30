package route

import (
	"apirestecommerce/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CategoryRoutes(app *fiber.App,db *gorm.DB) {
	h := &handler.CategoryHandler{
		DB: db,
	}
	category := app.Group("/categories")
	category.Get("/",h.Index)
	category.Get("/:id",h.Show)
	category.Post("/",h.Store)
	category.Delete("/:id",h.Delete)
}
