package jobs

import (
	"testing"
	"time"

	"github.com/arschles/assert"
)

type testPeriodic struct {
	t    *testing.T
	err  error
	i    int
	freq time.Duration
}

func (t *testPeriodic) Do() error {
	t.t.Logf("testPeriodic Do at %s", time.Now())
	t.i++
	return t.err
}

func (t testPeriodic) Frequency() time.Duration {
	return t.freq
}

func TestDoPeriodic(t *testing.T) {
	interval := time.Duration(3000) * time.Millisecond
	p := &testPeriodic{t: t, err: nil, freq: interval}
	closeCh1 := DoPeriodic([]Periodic{p})
	time.Sleep(interval / 2) // wait a little while for the goroutine to call the job once
	assert.True(t, p.i == 1, "the periodic wasn't called once")
	time.Sleep(interval)
	assert.True(t, p.i == 2, "the periodic wasn't called twice")
	time.Sleep(interval)
	assert.True(t, p.i == 3, "the periodic wasn't called thrice")
	close(closeCh1)
}
