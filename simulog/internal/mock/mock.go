package mock

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/bruteforce"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/common"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/ddos"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/normal"
)

type Config = common.Config

func Run(ctx context.Context, config *Config) {
	for {
		mode, err := gofakeit.Weighted(
			[]any{
				normal.Mode,
				bruteforce.Mode,
				ddos.Mode,
			},
			[]float32{
				config.Normal,
				config.Bruteforce,
				config.DDoS,
			},
		)

		if err != nil {
			log.Error().Err(err).Msg("Failed to randomly select mode")
			continue
		}

		done := make(chan bool, 1)

		switch mode {
		case normal.Mode:
			go normal.Run(config, done)
		case bruteforce.Mode:
			go bruteforce.Run(config, done)
		case ddos.Mode:
			go ddos.Run(config, done)
		default:
			log.Error().Msgf("Unknown mode `%v`", mode)
		}

		select {
		case <-done:
			continue
		case <-ctx.Done():
			return
		}
	}
}
