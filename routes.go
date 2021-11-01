package main

import (
	"net/http"

	"emailscraper/handlers"
	"emailscraper/models"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// CreateRoutes initialize all routes
func createRoutes(app models.RouterApp) {

	app.Get(
		"/ping", func(ctx *fiber.Ctx) error {
			log.Info().Msg("Testing logger")
			return ctx.SendStatus(http.StatusOK)
		},
	)

	app.Post("/scraper", handlers.ScrapeEmails)

}
