package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func BadRequestResponse(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusBadRequest).JSON(WebResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: err.Error(),
	})
}

func InternalServerErrorResponse(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(WebResponse{
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: err.Error(),
	})
}

func NotFoundResponse(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusNotFound).JSON(WebResponse{
		Status:  http.StatusText(http.StatusNotFound),
		Message: err.Error(),
	})
}
