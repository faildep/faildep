package faildep

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMetric_add_and_get(t *testing.T) {

	n := Resource{
		index:  1,
		Server: "123",
	}

	m := newNodeMetric(func() (func() ResourceList, chan struct{}) {
		return func() ResourceList {
			return ResourceList{n}
		}, make(chan struct{})
	})

	nm := m.takeMetric(n)
	nm.recordFailure(1 * time.Millisecond)
	m.takeMetric(n).recordFailure(1 * time.Millisecond)
	mm := m.takeMetric(n)
	assert.Equal(t, uint64(2), mm.takeFailCount())

	m.takeMetric(n).recordSuccess(1 * time.Millisecond)

	mm = m.takeMetric(n)

	assert.Equal(t, uint64(0), mm.takeFailCount())

}
