package throttle

import (
	"testing"
	"time"
)

var (
	mockStore = make(map[string]Limiter)
	window    = time.Second
	limit     = 10
	ipAddress = "192.168.1.1"
)

func Test_throttle(t *testing.T) {
	handler := New(mockStore, window, limit)
	handler.throttle(ipAddress)
}

func Test_throttle_bulk(t *testing.T) {
	var got bool
	for i := 0; i < 100; i++ {
		handler := New(mockStore, window, limit)
		got = handler.throttle(ipAddress)
	}

	if got != false {
		t.Errorf("throttle() = %v, want %v", got, false)
	}

}
