package service

import (
	"encoding/json"
	"fmt"
	"go-music-scrapper/router"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Pokemon struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location_area_encounters"`
}

type User struct{}

func BuildUserService(v1 *router.V1) {
	v1.Get("/user", authUser)
	v1.Get("/user/playlist", userPlaylist)
	v1.Get("/pokemon", getPokemon)
}

func authUser(c *fiber.Ctx) error {
	return c.SendString("Hello User")
}

func userPlaylist(c *fiber.Ctx) error {
	return c.SendString("user playlist")
}

// Example of consuming API
func getPokemon(c *fiber.Ctx) error {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		c.SendString(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.SendString(err.Error())
	}
	var poke Pokemon
	json.Unmarshal(body, &poke)

	msg := fmt.Sprintf("Name %v with the Id of %d, You can find it here %v", poke.Name, poke.Id, poke.Location)

	return c.SendString(msg)
}
