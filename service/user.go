package service

import (
	"go-music-scrapper/router"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id int
}

func BuildUserService(v1 *router.V1) {
	v1.Get("/user", authUser)
	v1.Get("/artist", getArtist)
}

func authUser(c *fiber.Ctx) error {
	return c.SendString("Hello User")
}

func getArtist(c *fiber.Ctx) error {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		c.SendString(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.SendString(err.Error())
	}

	return c.SendString(string(body))
}
