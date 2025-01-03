package normal

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/common"
)

const Mode = "normal"

func Run(config *common.Config, done chan<- bool) {
	method, err := gofakeit.Weighted(
		common.Methods,
		[]float32{
			0.7,
			0.2,
			0.07,
			0.03,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to randomly select HTTP method")
	}

	var endpoint string

	switch method {
	case http.MethodGet:
		endpoint = gofakeit.RandomString(common.GetEndpoints)
	case http.MethodPost:
		endpoint = gofakeit.RandomString(common.PostEndpoints)
	case http.MethodPut:
		endpoint = gofakeit.RandomString(common.PutEndpoints)
	case http.MethodDelete:
		endpoint = gofakeit.RandomString(common.DeleteEndpoints)
	}

	status, err := gofakeit.Weighted(
		common.Status,
		[]float32{
			0.7,
			0.1,
			0.05,
			0.04,
			0.03,
			0.04,
			0.04,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to randomly select HTTP status")
	}

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

	log.Info().
		Str("id", gofakeit.UUID()).
		Str("method", method.(string)).
		Str("endpoint", endpoint).
		Any("status", status).
		Str("ip", gofakeit.IPv4Address()).
		Str("ua", gofakeit.UserAgent()).
		Uint("ms", responseTime).
		Send()

	common.Sleep(config.T)
	done <- true
}
