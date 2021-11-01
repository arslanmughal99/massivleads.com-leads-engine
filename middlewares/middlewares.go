package middlewares

import (
	smodels "emailscraper/models"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetMiddlewares set middlewares on http app instance
func SetMiddlewares(app smodels.RouterApp) {
	auth := BasicAuth()

	app.Use(auth)
	app.Use(recover.New())
	app.Use(compress.New())
}
