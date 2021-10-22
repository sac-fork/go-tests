package go_tests

import (
	"testing"
	"time"
)

func TestSerial(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		println("Running First test =====")
		time.Sleep(time.Second)
		println("===== First test complete")
	})

	t.Run("Second", func(t *testing.T) {
		println("Running Second test =====")
		time.Sleep(time.Second)
		println("===== Second test complete")
	})

	t.Run("Third", func(t *testing.T) {
		println("Running Third test =====")
		time.Sleep(time.Second)
		println("===== Third test complete")
	})

	t.Run("Fourth", func(t *testing.T) {
		println("Running Fourth test =====")
		time.Sleep(time.Second)
		println("===== Fourth test complete")
	})

	t.Run("Fifth", func(t *testing.T) {
		println("Running Fifth test =====")
		time.Sleep(time.Second)
		println("===== Fifth test complete")
	})
}