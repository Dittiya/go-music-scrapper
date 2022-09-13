package server

import (
	"fmt"

	"github.com/Dittiya/go-music-scrapper/config"
	"github.com/gofiber/fiber/v2"
)

func NewServer() {
	app := fiber.New()
	cfg := config.NewConfig()

	app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
