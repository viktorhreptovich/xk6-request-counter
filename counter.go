package counter

import (
	"sync/atomic"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/counter", New())
}

var requestCounter int64
var successCounter int64
var failureCounter int64

type counter struct{}

func (c *counter) Success() int64 {
	atomic.AddInt64(&requestCounter, 1)
	return atomic.AddInt64(&successCounter, 1)
}

func (c *counter) Fail() int64 {
	atomic.AddInt64(&requestCounter, 1)
	return atomic.AddInt64(&failureCounter, 1)
}

func New() *counter {
	return &counter{}
}
