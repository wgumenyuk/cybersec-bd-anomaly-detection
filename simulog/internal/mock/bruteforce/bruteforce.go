package bruteforce

import "github.com/wgumenyuk/cybersec-bd-anomaly-detection/simulog/internal/mock/common"

const Mode = "bruteforce"

func Run(config *common.Config, done chan<- bool) {
	common.Sleep(config.T)
	done <- true
}
