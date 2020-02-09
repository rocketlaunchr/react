// +build go1.13

package context

import (
	"context"
	"time"
)

var Canceled = context.Canceled

var DeadlineExceeded error = context.DeadlineExceeded

func WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(parent)
}

func WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(parent, d)
}

func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

func Background() context.Context {
	return context.Background()
}

func TODO() context.Context {
	return context.TODO()
}

func WithValue(parent context.Context, key, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}
