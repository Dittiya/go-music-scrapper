package server

import (
	"fmt"
	"go-music-scrapper/config"
	"go-music-scrapper/db"
	"go-music-scrapper/service"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

func NewServer() {
	cfg := config.DefaultConfig()
	app := fiber.New(cfg.Config)

	app.Use(cors.New())

	dbase := db.Redis{
		Options: &redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		},
	}

	dbase.InitRedis()

	v1 := service.InjectUserService(app)
	service.BuildUserService(v1, &dbase)

	app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
