package common

import (
	"math/rand/v2"
	"time"
)

func Sleep(t uint) {
	time.Sleep(time.Duration(rand.UintN(t)+1) * time.Second)
}
