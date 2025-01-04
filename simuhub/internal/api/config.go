package api

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simuhub/internal/etcd"
)

type Config struct {
	T          uint    `json:"t"`
	Normal     float32 `json:"normal"`
	Bruteforce float32 `json:"bruteforce"`
	DDoS       float32 `json:"ddos"`
}

func HandleGetConfig(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response, err := etcd.Client.Get(ctx, etcd.Key)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get config from Etcd")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(response.Kvs) == 0 {
		return c.Status(fiber.StatusOK).JSON(&Config{
			T:          5,
			Normal:     1,
			Bruteforce: 0,
			DDoS:       0,
		})
	}

	config := new(Config)

	if err := sonic.Unmarshal(response.Kvs[0].Value, config); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal config")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(config)
}

func HandleUpdateConfig(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	body := new(Config)

	if err := c.Bind().Body(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	config, err := sonic.MarshalString(body)

	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal config")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if _, err := etcd.Client.Put(ctx, etcd.Key, config); err != nil {
		log.Error().Err(err).Msg("Failed to write config to Etdc")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info().Msg("Updating config")

	return c.SendStatus(fiber.StatusOK)
}
