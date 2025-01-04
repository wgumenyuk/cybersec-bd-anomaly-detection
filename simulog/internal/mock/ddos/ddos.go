package ddos

import (
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	"github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/common"
)

const Mode = "ddos"

const (
	numberOfAttackers = 5
	numberOfEndpoints = 20
	baseResponseTime  = 1500
	ticks             = 10
)

func Run(config *common.Config, done chan<- bool) {
	ips := make([]string, 0, numberOfAttackers)

	for i := 0; i < numberOfAttackers; i++ {
		ips = append(ips, gofakeit.IPv4Address())
	}

	endpoints := make([]string, 0, numberOfEndpoints)

	for i := 0; i < numberOfEndpoints; i++ {
		endpoints = append(
			endpoints,
			fmt.Sprintf("/api/v1/search?q=%s", gofakeit.Word()),
		)
	}

	ua := gofakeit.UserAgent()

	for i := 0; i < ticks; i++ {
		log.Info().
			Str("id", gofakeit.UUID()).
			Str("method", http.MethodGet).
			Str("endpoint", gofakeit.RandomString(endpoints)).
			Int("status", http.StatusOK).
			Str("ip", gofakeit.RandomString(ips)).
			Str("ua", ua).
			Uint("ms", baseResponseTime+100*(ticks+1)).
			Send()

		common.Sleep(config.T)
	}

	done <- true
}
