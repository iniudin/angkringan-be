package controller

import (
	"angkringan/api/model/request"
	"angkringan/api/model/response"
	"angkringan/pkg/entity"
	"angkringan/pkg/product"
	"angkringan/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ProductController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindByName(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	validator *validation.CustomValidator
	service   product.Service
}

func NewProductController(validator *validation.CustomValidator, service product.Service) *ProductControllerImpl {
	return &ProductControllerImpl{validator: validator, service: service}
}

func (c *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	body := request.CreateProduct{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    err.Error(),
		})
	}

	if err := c.validator.Validate(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Some field should be provided",
			Data:    err,
		})
	}
	newProduct, err := c.service.Create(ctx.Context(), body)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success create a new product",
		Data:    newProduct,
	})
}

func (c *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	body := entity.Product{}
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data:    body,
		},
	)
}

func (c *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	body := entity.Product{}
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data:    body,
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
func (c *ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	pageNumber, _ := strconv.Atoi(ctx.Query("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.Query("size", "10"))

	products, err := c.service.FindAll(ctx.Context(), pageNumber, pageSize)
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

func (c *ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	body := entity.Product{}
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data:    body,
		},
	)
}

func (c *ProductControllerImpl) FindByName(ctx *fiber.Ctx) error {
	body := entity.Product{}
	return ctx.JSON(
		response.WebResponse{
			Status:  "success",
			Message: "get all product",
			Data:    body,
		},
	)
}
