package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/etcd"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Logger = zerolog.
		New(os.Stdout).
		With().
		Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if _, ok := os.LookupEnv("DEBUG"); ok {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if err := etcd.Connect(); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Etcd")
	}

	defer etcd.Client.Close()

	config := &mock.Config{
		T:          5,
		Normal:     1,
		DDoS:       0,
		Bruteforce: 0,
	}

	go func() {
		ch := etcd.Client.Watch(context.Background(), etcd.Key)

		for response := range ch {
			for _, event := range response.Events {
				if err := sonic.Unmarshal(event.Kv.Value, config); err != nil {
					log.Debug().Err(err).Msg("Failed to unmarshal config")
					continue
				}

				log.Debug().Msg("Updating config")
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go mock.Run(ctx, config)

	<-quit
}
