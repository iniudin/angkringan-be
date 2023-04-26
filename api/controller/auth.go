package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
}

func NewAuthController() *AuthControllerImpl {
	return &AuthControllerImpl{}
}

func (c *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c *AuthControllerImpl) Refresh(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
