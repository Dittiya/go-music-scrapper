package config

import "github.com/gofiber/fiber/v2"

type Config struct {
	fiber.Config
	Port int
}

func DefaultConfig() Config {
	cfg := fiber.Config{AppName: "Music Scrapper"}
	port := 8008
	return Config{Config: cfg, Port: port}
}
