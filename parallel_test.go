package go_tests

import (
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		t.Parallel()
		println("Running First test =====")
		time.Sleep(time.Second)
		println("===== First test complete")
	})

	t.Run("Second", func(t *testing.T) {
		t.Parallel()
		println("Running Second test =====")
		time.Sleep(time.Second)
		println("===== Second test complete")
	})

	t.Run("Third", func(t *testing.T) {
		t.Parallel()
		println("Running Third test =====")
		time.Sleep(time.Second)
		println("===== Third test complete")
	})

	t.Run("Fourth", func(t *testing.T) {
		t.Parallel()
		println("Running Fourth test =====")
		time.Sleep(time.Second)
		println("===== Fourth test complete")
	})

	t.Run("Fifth", func(t *testing.T) {
		t.Parallel()
		println("Running Fifth test =====")
		time.Sleep(time.Second)
		println("===== Fifth test complete")
	})
}