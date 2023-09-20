package server

import (
	"os"

	"discord-oauth2/internal/app/server/routes/oauth2"
	"discord-oauth2/internal/app/server/routes/user"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() error {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.ConfigFastest.Marshal,
		JSONDecoder: sonic.ConfigFastest.Unmarshal,
	})
	app.Use(cors.New())

	api := app.Group("/api")

	oauth2Endpoint := api.Group("/oauth2")
	oauth2Endpoint.Get("/url", oauth2.GetURL)
	oauth2Endpoint.Get("/token", oauth2.GetToken)

	userEndpoint := api.Group("/user")
	userEndpoint.Get("/", user.GetMe)

	return app.Listen(":" + os.Getenv("PORT"))
}
