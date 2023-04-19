package route

import (
	"angkringan/api/controller"
	"angkringan/pkg/product"
	"angkringan/pkg/validation"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewProductRoute(router fiber.Router, db *sql.DB, validator *validation.CustomValidator) {
	repository := product.NewRepository(db)
	service := product.NewService(repository)
	endpoint := controller.NewProductController(validator, service)

	products := router.Group("/products")
	products.Post("", endpoint.Create)
	products.Put("", endpoint.Update)
	products.Delete("", endpoint.Delete)
	products.Get("", endpoint.FindAll)
	products.Get("/", endpoint.FindById)

}
