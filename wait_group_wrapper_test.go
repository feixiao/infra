package infra

import (
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg WaitGroupWrapper
	wg.Wrap(func() {
		time.Sleep(2 * time.Millisecond)
	})

	wg.Wait()
}
