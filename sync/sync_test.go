package sync

import (
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {
	t.Run("Increment counter 3 times in 1 thread", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
				
		assertCount(t, &counter, 3)
	})

	t.Run("Runs safely in a concurrent env", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i :=0; i< wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()
		assertCount(t, &counter, wantedCount)

	})
}


func assertCount(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got, want)
	}
}