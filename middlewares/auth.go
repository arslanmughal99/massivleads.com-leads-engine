package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func BasicAuth() fiber.Handler {
	return basicauth.New(
		basicauth.Config{
			Users: map[string]string{
				os.Getenv("BASICAUTH_USERNAME"): os.Getenv("BASICAUTH_PASSWORD"),
			},
		},
	)
}
