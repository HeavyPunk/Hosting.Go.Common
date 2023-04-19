package test_retry

import (
	"fmt"
	tools_retry "simple-hosting/go-common/tools/retry"
	"testing"
	"time"
)

func TestRetry5Attempts(t *testing.T) {
	counter := 0
	delay, _ := time.ParseDuration("0s")
	tools_retry.Retry(
		func() (bool, error) {
			fmt.Println("Retrying...")
			counter++
			return true, nil
		},
		func(bool, error) bool { return counter == 5 },
		10,
		delay,
		nil,
	)
	if counter != 5 {
		t.Errorf("Expected 5 retries, got %v", counter)
	}
}
