package account

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

func (mw loggingMiddleware) CreateUser(ctx context.Context, email string, password string) (resp string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateUser", "email", email, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.CreateUser(ctx, email, password)
}

func (mw loggingMiddleware) GetUser(ctx context.Context, id string) (resp string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetUser", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetUser(ctx, id)
}
