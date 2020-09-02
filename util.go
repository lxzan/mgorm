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
func Context() context.Context {
	return WithTimeout(15 * time.Second)
}
