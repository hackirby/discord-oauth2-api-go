package oauth2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"os"
)

func GetURL(c *fiber.Ctx) error {
	OAuth2 := oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}
	return c.JSON(fiber.Map{
		"url": OAuth2.AuthCodeURL("state"),
	})
}
