package main

import (
	"cashflow/internal/app/server"
	"cashflow/internal/pkg/config"
	"cashflow/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.New()
	db := database.New(cfg)
	fiberApp := fiber.New()

	app := server.New(db, fiberApp)

	app.Run()
}
