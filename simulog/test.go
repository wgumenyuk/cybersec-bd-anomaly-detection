package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	var (
		a	float32 = 1
		b	float32 = 0
	)

	done := make(chan bool, 1)

	go func(x *float32, y *float32) {
		for {
			v, _ := gofakeit.Weighted(
				[]any{
					"A",
					"B",
				},
				[]float32{
					*x,
					*y,
				},
			)

			fmt.Println(v, *x, *y)
			time.Sleep(time.Second)
		}
	}(&a, &b)

	time.Sleep(2*time.Second)

	a = 0.5
	b = 0.5

	<-done
}
