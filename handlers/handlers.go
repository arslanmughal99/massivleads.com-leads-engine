package handlers

import (
	"emailscraper/models"
	"github.com/gofiber/fiber/v2"
)

// SendResponse is a generic method to send http response using current http framework
func SendResponse(ctx models.RouterCtx, res *models.Result) error {
	if res.Exception != nil {
		return ctx.Status(int(res.Exception.StatusCode)).JSON(res.Exception)
	}

	return ctx.Status(fiber.StatusOK).JSON(res.Result)
}
