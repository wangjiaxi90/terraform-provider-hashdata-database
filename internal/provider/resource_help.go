package provider

import (
	"math/rand"
	"time"
)

type stop struct {
	error
}

func simpleRetry(fn func() error) error {
	return retry(100, 10*time.Second, fn)
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			//logger.Warn("Retry function")
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}

	return nil
}
