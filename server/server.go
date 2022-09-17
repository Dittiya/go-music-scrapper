package server

import (
	"fmt"
	"go-music-scrapper/config"
	"go-music-scrapper/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer() {
	cfg := config.DefaultConfig()
	app := fiber.New(cfg.Config)

	app.Use(cors.New())

	v1 := service.InjectUserService(app)
	service.BuildUserService(v1)

	app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
