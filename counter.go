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
var okCounter int64
var errorCounter int64

type counter struct{}

func (c *counter) UpOk() int64 {
	atomic.AddInt64(&reqCounter, 1)
	return atomic.AddInt64(&okCounter, 1)

}
func (c *counter) UpError() int64 {
	atomic.AddInt64(&reqCounter, 1)
	return atomic.AddInt64(&errorCounter, 1)
}

func (c *counter) getOkCount() int64 {
	return atomic.LoadInt64(&okCounter)
}

func (c *counter) getErrorCount() int64 {
	return atomic.LoadInt64(&errorCounter)
}

func (c *counter) okRate() int64 {
	return atomic.LoadInt64(&okCounter) * 100 / atomic.LoadInt64(&reqCounter)
}

func (c *counter) errorRate() int64 {
	return atomic.LoadInt64(&errorCounter) * 100 / atomic.LoadInt64(&reqCounter)
}

func New() *counter {
	return &counter{}
}
