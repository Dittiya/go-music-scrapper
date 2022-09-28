package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-music-scrapper/db"
	"go-music-scrapper/router"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Pokemon struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location_area_encounters"`
}

type Token struct {
	AcessToken   string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type State string

func BuildUserService(v1 *router.V1, storage db.Storage) {
	v1.Get("/user", authUser)
	v1.Get("/user/playlist", userPlaylist)
	v1.Get("/login", spotifyLogin)
	v1.Get("/callback", spotifyCallback(storage))
	v1.Get("/pokemon", getPokemon)
}

func authUser(c *fiber.Ctx) error {
	return c.SendString("Yooo")
}

// TODO
// Remove hardcoded redirect_uri
// Improve this, that format is hideous
func spotifyLogin(c *fiber.Ctx) error {
	url := "https://accounts.spotify.com/authorize?"
	respType := "code"
	clientId := os.Getenv("CLIENT_ID")
	scope := "user-read-private user-read-email"
	redirectUri := "http://localhost:8008/api/v1/callback"
	state := "testtesttesttest" // Create function to randomize this

	uri := fmt.Sprintf(
		"%vresponse_type=%v&client_id=%v&scope=%v&redirect_uri=%v&state=%v",
		url, respType, clientId, scope, redirectUri, state,
	)

	return c.Redirect(uri)
}

// TODO
// Cache Access Token according to expiring time
// Remove hardcoded redirect_uri
func spotifyCallback(store db.Storage) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bodyParams := url.Values{
			"grant_type":   {"authorization_code"},
			"code":         {c.Query("code", "empty")},
			"redirect_uri": {"http://localhost:8008/api/v1/callback"},
		}

		url := "https://accounts.spotify.com/api/token"
		bearer := "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID")+":"+os.Getenv("CLIENT_SECRET")))
		req, _ := http.NewRequest("POST", url, strings.NewReader(bodyParams.Encode()))
		req.Header.Add("Authorization", bearer)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.SendString(err.Error())
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.SendString(err.Error())
		}

		var token Token
		json.Unmarshal(body, &token)

		return c.SendString(string(body))
	}
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
