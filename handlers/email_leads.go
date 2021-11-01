package handlers

import (
	"net/http"
	"strings"

	"emailscraper/dtos"
	"emailscraper/exceptions"
	"emailscraper/models"
	"emailscraper/services"
	"emailscraper/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// ScrapeEmails scrape emails and post data to webhook
func ScrapeEmails(ctx models.RouterCtx) error {
	dto := new(dtos.EmailScraper)

	if err := ctx.BodyParser(dto); err != nil {
		log.Error().Err(errors.Wrap(err, "handlers.ScrapeEmails")).Msg("Failed to parse body")
		return err
	}

	if err := utils.ValidateDto(dto); err != nil {
		log.Debug().Interface("dto", *dto).Strs("reasons", *err).Msg("handlers.ScrapeEmails Dto validation failed")

		res := new(exceptions.BaseException)
		res.Error = "Bad Request"
		res.StatusCode = fiber.StatusBadRequest
		res.Message = strings.Join(*err, "\n")

		return ctx.Status(fiber.StatusBadRequest).JSON(res)
	}

	go func() {
		_ = services.ScrapeEmailLeads(*dto)
	}()

	resp := new(dtos.EmailScraperResp)
	resp.Id = dto.Id

	return ctx.Status(http.StatusOK).JSON(resp)
}
