package network

import (
	"time"

	circuit "github.com/rubyist/circuitbreaker"
)

func GetHttpClient() (*circuit.HTTPClient, *circuit.Breaker) {
	circuitBreaker := circuit.NewThresholdBreaker(3)
	client := circuit.NewHTTPClientWithBreaker(circuitBreaker, time.Second*1, nil)
	return client, circuitBreaker
}
