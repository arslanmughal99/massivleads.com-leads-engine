package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"emailscraper/exceptions"
	"emailscraper/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func main() {
	initLogger()

	procSize, err := strconv.ParseInt(os.Getenv("PROCS"), 10, 0)
	if err != nil {
		procSize = int64(runtime.GOMAXPROCS(-1))
	}

	runtime.GOMAXPROCS(int(procSize))

	app := fiber.New(
		fiber.Config{
			CaseSensitive: true,
			ServerHeader:  os.Getenv("SERVER_HEADER"),
			Prefork:       os.Getenv("PREFORK") == "yes",
			ErrorHandler:  exceptions.GlobalExceptionHandler,
		},
	)

	// Initialize middlewares
	middlewares.SetMiddlewares(app)

	// Initialize routes
	createRoutes(app)

	log.Fatal().Err(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
