package user

import (
	"discord-oauth2/internal/app/database"
	"discord-oauth2/internal/app/database/models"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"os"
)

func GetMe(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing token",
		})
	}

	OAuth2 := oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	client := OAuth2.Client(c.Context(), &oauth2.Token{
		AccessToken: token,
	})

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
			"error": "failed to decode user json",
		})
	}

	dbUser := models.User{}
	err = database.Database.First(&dbUser, models.User{
		UserID: user.ID,
	}).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"user": dbUser,
	})
}
