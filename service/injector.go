//go:build wireinject
// +build wireinject

package service

import (
	"go-music-scrapper/router"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InjectUserService(app *fiber.App) *router.V1 {
	wire.Build(router.BuildApiEndpoint, router.BuildV1Endpoint)
	return &router.V1{}
}
