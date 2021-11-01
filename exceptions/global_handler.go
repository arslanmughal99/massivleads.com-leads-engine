package exceptions

import (
	"github.com/gofiber/fiber/v2"
	serrors "github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// GlobalExceptionHandler handle any global exception
func GlobalExceptionHandler(ctx *fiber.Ctx, err error) error {
	res := new(BaseException)
	res.Error = "Internal Server Error"
	res.Message = "Something went wrong"
	res.StatusCode = fiber.StatusInternalServerError

	log.Error().Err(serrors.Wrap(err, "exceptions.GlobalExceptionHandler")).Msg("Global error occur")

	return ctx.Status(int(res.StatusCode)).JSON(res)
}
