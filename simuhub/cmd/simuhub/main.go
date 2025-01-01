package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simuhub/internal/api"
)

const port = ":8080"

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use("/", static.New("static"))

	app.Get("/api/v1/status", api.HandleGetStatus)
	app.Get("/api/v1/config", api.HandleGetConfig)
	app.Put("/api/v1/status", api.HandleUpdateConfig)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		app.Listen(port, fiber.ListenConfig{
			DisableStartupMessage: true,
		})
	}()

	log.Info().Msgf("Listening on port %s", port)

	<-quit
	log.Info().Msg("Exiting")

	app.Shutdown()
}
