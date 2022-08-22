package throttle

import "time"

type throttleHandler struct {
	store  map[string]Limiter
	window time.Duration
	limit  int
}

type ThrottleHandler interface {
	throttle(ipAddress string) bool
}

type Limiter struct {
	// stored in unixnano.
	Timestamp int

	RequestCount int
}

func New(store map[string]Limiter, window time.Duration, limit int) ThrottleHandler {
	return &throttleHandler{
		store:  store,
		window: window,
		limit:  limit,
	}
}

func (h *throttleHandler) throttle(ipAddress string) bool {
	if _, ok := h.store[ipAddress]; !ok {
		h.store[ipAddress] = Limiter{
			Timestamp:    int(time.Now().UnixNano()),
			RequestCount: 1,
		}
		return true
	} else if h.store[ipAddress].RequestCount > h.limit {
		return false
	}

	return true
}
