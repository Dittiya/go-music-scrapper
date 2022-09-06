package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := ":3030"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s", c.Params("*"))
		return c.SendString(msg)
	}).Name("api")

	data, _ := json.MarshalIndent(app.GetRoute("api"), "", " ")
	fmt.Print(string(data))

	app.Listen(port)
}
