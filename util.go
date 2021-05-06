package mgorm

import (
	"context"
	"time"
)

func WithTimeout(d time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), d)
	return ctx
}

// default context
func NewContext() context.Context {
	return WithTimeout(10 * time.Second)
}

func WithWrap(ctx context.Context) context.Context {
	if ctx == nil {
		return NewContext()
	}
	return ctx
}
