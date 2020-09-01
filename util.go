package mgorm

import (
	"context"
	"time"
)

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx
}
