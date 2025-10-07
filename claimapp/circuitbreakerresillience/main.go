package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	"golang.org/x/time/rate"
)

// RateLimitMiddleware limits executions to the limiter's rate.
func RateLimitMiddleware(l *rate.Limiter) goresilience.Middleware {
	return func(next goresilience.Runner) goresilience.Runner {
		return goresilience.RunnerFunc(func(ctx context.Context, f goresilience.Func) error {
			if !l.Allow() {
				return fmt.Errorf("rate limit exceeded")
			}
			//next = goresilience.Sanitize(next)
			return next.Run(ctx, f)
		})
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Rate: 1 req/s with burst 5
	rlMw := RateLimitMiddleware(rate.NewLimiter(rate.Limit(1), 5)) // 1 req/sec, burst 5 (very lenient)

	cbMw := circuitbreaker.NewMiddleware(circuitbreaker.Config{
		ErrorPercentThresholdToOpen:        50,
		MinimumRequestToOpen:               4,
		WaitDurationInOpenState:            5 * time.Second,
		MetricsSlidingWindowBucketQuantity: 5,
		MetricsBucketDuration:              1 * time.Second,
	})

	retryMw := retry.NewMiddleware(retry.Config{
		Times:          5,
		WaitBase:       300 * time.Millisecond,
		DisableBackoff: false, // keep built-in exponential backoff+jitter
	})

	timeoutMw := timeout.NewMiddleware(timeout.Config{Timeout: 2 * time.Second})

	// Order: RateLimit → Timeout → CircuitBreaker → Retry
	runner := goresilience.RunnerChain(rlMw, timeoutMw, cbMw, retryMw)

	ctx := context.Background()
	//for i := 1; i <= 12; i++ {
	//log.Printf("\n=== Attempt %d ===", i)
	err := runner.Run(ctx, func(ctx context.Context) error {
		resp, err := http.Get("https://jsonplaceholder.typcode.com/posts/1")
		if err != nil {
			log.Printf("HTTP error: %v", err)
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status %d", resp.StatusCode)
		}
		b, _ := io.ReadAll(resp.Body)
		log.Printf("OK (%d bytes)", len(b))
		return nil
	})
	if err != nil {
		log.Printf("Run() error: %v", err)
	}
	time.Sleep(100 * time.Millisecond) // just to make logs readable
	//}
}
