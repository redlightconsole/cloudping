package cloudping

import (
	"sync"
)

// A Runner waits for a collection of goroutines to finish.
// It's wrapper around the waitgroup struct to add concurrency limit.
type Runner struct {
	ch chan int
	wg *sync.WaitGroup
}

// NewRunner creates a new runner for concurrency tasks
func NewRunner(limit int) *Runner {
	if limit <= 0 {
		limit = 1
	}
	return &Runner{
		ch: make(chan int, limit), // buffer chan to limit concurrency
		wg: &sync.WaitGroup{},
	}
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// See sync.WaitGroup for more info.
func (wgl *Runner) Add(delta int) {
	for i := 0; i < delta; i++ {
		wgl.ch <- 1
		wgl.wg.Add(1)
	}
}

// Done decrements the WaitGroup counter by one.
func (wgl *Runner) Done() {
	wgl.wg.Done()
	<-wgl.ch
}

// Wait blocks until the WaitGroup counter is zero.
func (wgl *Runner) Wait() {
	close(wgl.ch)
	wgl.wg.Wait()
}
