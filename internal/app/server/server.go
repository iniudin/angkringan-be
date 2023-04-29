package server

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Run()
}

type ServerImpl struct {
	*sql.DB
	*fiber.App
}

func New(db *sql.DB, app *fiber.App) *ServerImpl {
	return &ServerImpl{
		DB:  db,
		App: app,
	}
}

func (s *ServerImpl) Run() {
	go func() {
		if err := s.App.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = s.App.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	if err := s.DB.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fiber was successful shutdown.")
}
