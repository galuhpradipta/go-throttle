package throttle

import "time"

type throttleHandler struct {
	store  map[string]Limiter
	window time.Duration
}

type ThrottleHandler interface{}

type Limiter struct {
	// stored in unixnano.
	Timestamp int

	Limit int
}

func New(window time.Duration, store map[string]Limiter) ThrottleHandler {
	return &throttleHandler{
		store:  store,
		window: window,
	}
}

func throttle() bool {
	return true
}
