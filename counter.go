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

func (c *counter) SuccessCount() int64 {
	return atomic.LoadInt64(&successCounter)
}

func (c *counter) FailureCount() int64 {
	return atomic.LoadInt64(&failureCounter)
}

func (c *counter) successRate() int64 {
	return (atomic.LoadInt64(&successCounter) * 100) / atomic.LoadInt64(&requestCounter)
}

func (c *counter) FailureRate() int64 {
	return (atomic.LoadInt64(&failureCounter) * 100) / atomic.LoadInt64(&requestCounter)
}

func New() *counter {
	return &counter{}
}
