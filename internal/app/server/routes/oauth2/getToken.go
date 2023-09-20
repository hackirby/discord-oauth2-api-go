package oauth2

import (
	"discord-oauth2/internal/app/database"
	"discord-oauth2/internal/app/database/models"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"os"
)

func GetToken(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "code is required",
		})
	}

	OAuth2 := oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	token, err := OAuth2.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid code",
		})
	}

	client := OAuth2.Client(c.Context(), token)
	res, err := client.Get("https://discord.com/api/v9/users/@me")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	if res.StatusCode != fiber.StatusOK {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	type User struct {
		ID string `json:"id"`
	}

	var user User
	err = sonic.ConfigFastest.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to decode user",
		})
	}

	var dbUser models.User
	err = database.Database.FirstOrCreate(&dbUser, models.User{
		UserID: user.ID,
	}).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to find or create user",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
