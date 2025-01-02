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
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simuhub/internal/etcd"
)

const port = ":8080"

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	defer app.Shutdown()

	app.Use("/", static.New("static"))

	app.Get("/api/v1/status", api.HandleGetStatus)
	app.Get("/api/v1/config", api.HandleGetConfig)
	app.Put("/api/v1/config", api.HandleUpdateConfig)

	if err := etcd.Connect(); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Etcd")
	}

	defer etcd.Client.Close()

	go func() {
		app.Listen(port, fiber.ListenConfig{
			DisableStartupMessage: true,
		})
	}()

	log.Info().Msgf("Listening on port %s", port)

	<-quit
}
