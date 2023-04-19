package controller

import (
	"angkringan/api/model/response"
	"angkringan/pkg/product"
	"angkringan/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type ProductControllerImpl struct {
	validator *validation.CustomValidator
	service   product.Service
}

func NewProductController(validator *validation.CustomValidator, service product.Service) *ProductControllerImpl {
	return &ProductControllerImpl{validator: validator, service: service}
}

func (p *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "create a product",
			Data: response.ProductResponse{
				ID:          "",
				Name:        "",
				Description: "",
				Price:       0,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	)
}

func (p *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data: response.ProductResponse{
				ID:          "",
				Name:        "",
				Description: "",
				Price:       0,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	)
}

func (p *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data: response.ProductResponse{
				ID:          "",
				Name:        "",
				Description: "",
				Price:       0,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	)
}

// FindAll is a function to get all products data from database
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} response.WebResponse{data=[]response.ProductResponse}
// @Failure 503 {object} response.WebResponse{}
// @Router /api/v1/products [get]
func (p *ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	pageNumber, _ := strconv.Atoi(ctx.Query("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.Query("size", "10"))

	products, err := p.service.FindAll(ctx.Context(), pageNumber, pageSize)
	if err != nil {
		return ctx.Status(http.StatusServiceUnavailable).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusServiceUnavailable),
			Message: err.Error(),
			Data:    nil,
		})

	}
	return ctx.Status(http.StatusOK).JSON(
		response.WebResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "get all product",
			Data:    products,
		},
	)
}

func (p *ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data: response.ProductResponse{
				ID:          "",
				Name:        "",
				Description: "",
				Price:       0,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	)
}

func (p *ProductControllerImpl) FindByName(ctx *fiber.Ctx) error {
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data: response.ProductResponse{
				ID:          "",
				Name:        "",
				Description: "",
				Price:       0,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	)
}
