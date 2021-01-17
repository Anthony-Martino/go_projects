package auth

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

//LoggingMiddleware ...
type LoggingMiddleware func(Service) Service

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

//NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger) LoggingMiddleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw loggingMiddleware) Login(ctx context.Context, req LoginRequest) (resp string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Login", "request", req, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Login(ctx, req)
}

func (mw loggingMiddleware) Register(ctx context.Context, req RegisterRequest) (resp string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Register", "request", req, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.Register(ctx, req)
}

func (mw loggingMiddleware) GetUser(ctx context.Context, id string) (resp string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetUser", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetUser(ctx, id)
}
