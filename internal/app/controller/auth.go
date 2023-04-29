package controller

import (
	"cashflow/internal/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx fiber.Ctx) error
	Register(ctx fiber.Ctx) error
}

type AuthControllerImpl struct {
	Service user.Service
}

func NewAuthController(service user.Service) *AuthControllerImpl {
	return &AuthControllerImpl{Service: service}
}

func (c *AuthControllerImpl) Login(ctx fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c *AuthControllerImpl) Register(ctx fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
