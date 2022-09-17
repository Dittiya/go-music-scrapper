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

type User struct {
	Code  string
	State string
}

func BuildUserService(v1 *router.V1) {
	v1.Get("/user", authUser)
	v1.Get("/user/playlist", userPlaylist)
	v1.Get("/login", spotifyLogin)
	v1.Get("/callback", spotifyCallback)
	v1.Get("/pokemon", getPokemon)
}

func authUser(c *fiber.Ctx) error {
	return c.SendString("Yooo")
}

// TODO
// Improve this, that format is hideous
func spotifyLogin(c *fiber.Ctx) error {
	url := "https://accounts.spotify.com/authorize?"
	respType := "code"
	clientId := "035da001ba5b492ba5c527d149dc34e2"
	scope := "user-read-private user-read-email"
	redirectUri := "http://localhost:8008/api/v1/callback"
	state := "aAdf3i34O22LL19d"

	uri := fmt.Sprintf(
		"%vresponse_type=%v&client_id=%v&scope=%v&redirect_uri=%v&state=%v",
		url, respType, clientId, scope, redirectUri, state,
	)

	return c.Redirect(uri)
}

// TODO
// Save the Code and State then process get User's details
func spotifyCallback(c *fiber.Ctx) error {
	callback := User{
		Code:  c.Query("code", "empty"),
		State: c.Query("state", "empty"),
	}
	call, err := json.Marshal(callback)
	if err != nil {
		panic(err)
	}

	return c.SendString(string(call))
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
	po, err := json.Marshal(poke)
	if err != nil {
		panic(err)
	}

	return c.SendString(string(po))
}
