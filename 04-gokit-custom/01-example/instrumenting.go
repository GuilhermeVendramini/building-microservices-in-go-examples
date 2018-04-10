package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           UserService
}

func (mw instrumentingMiddleware) Get(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Get(s)
	return
}

func (mw instrumentingMiddleware) Delete(s int) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Delete(s)
	return
}
