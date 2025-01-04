package bruteforce

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/common"
)

const Mode = "bruteforce"

const ticks = 10

func Run(config *common.Config, done chan<- bool) {
	ip := gofakeit.IPv4Address()
	ua := gofakeit.UserAgent()

	responseTimeRange, err := gofakeit.Weighted(
		common.ResponseTimesRanges,
		[]float32{
			0.2,
			0.5,
			0.2,
			0.08,
			0.02,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to randomly select response time range")
	}

	var responseTime uint

	if v, ok := responseTimeRange.(common.ResponseTimeRange); ok {
		responseTime = gofakeit.UintRange(v.Min, v.Max)
	} else {
		log.Fatal().Err(err).Msg("Failed to typecast response time range")
	}

	for i := 0; i < ticks; i++ {
		log.Info().
			Str("id", gofakeit.UUID()).
			Str("method", http.MethodPost).
			Str("endpoint", "/api/v1/login").
			Any("status", http.StatusBadRequest).
			Str("ip", ip).
			Str("ua", ua).
			Uint("ms", responseTime).
			Send()

		common.Sleep(config.T)
	}

	done <- true
}
