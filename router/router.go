package router

import "github.com/gofiber/fiber/v2"

type API struct {
	fiber.Router
}

type V1 struct {
	fiber.Router
}

func BuildApiEndpoint(app *fiber.App) *API {
	api := app.Group("/api", func(c *fiber.Ctx) error {
		msg := "Using api group"
		c.Set("API-Notes", msg)

		return c.Next()
	})

	return &API{Router: api}
}

func BuildV1Endpoint(api *API) *V1 {
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		msg := "V1"
		c.Set("Version", msg)

		return c.Next()
	})

	return &V1{Router: v1}
}
