package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func (mw loggingMiddleware) Get(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Get(s)
	return
}

func (mw loggingMiddleware) Delete(s int) (n bool, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete",
			"input", s,
			"output", n,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.next.Delete(s)
	return
}
