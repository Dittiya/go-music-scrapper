package server

import (
	"fmt"
	"go-music-scrapper/config"
	"go-music-scrapper/service"

	"github.com/gofiber/fiber/v2"
)

func NewServer() {
	cfg := config.DefaultConfig()
	app := fiber.New(cfg.Config)

	v1 := service.InjectUserService(app)
	service.BuildUserService(v1)

	app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
