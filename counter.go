// This is a PoC/illustrative code to show how to share a single integer that goes up in k6 on a
// single instance

package counter

import (
	"sync/atomic"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/counter", New())
}

var reqCounter int64
var successCounter int64
var failedCounter int64

type counter struct{}

func (c *counter) Success() int64 {
	return atomic.AddInt64(&successCounter, 1)
}

func (c *counter) Fail() int64 {
	return atomic.AddInt64(&failedCounter, 1)
}

func New() *counter {
	return &counter{}
}
