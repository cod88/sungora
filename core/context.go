package core

import (
	"context"
	"time"
)

type Context struct {
	Data string
	context.Context
}

// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent context.Context, key, val interface{}) *Context {
	var ctx = new(Context)
	ctx.Context = context.WithValue(parent, key, val)
	return ctx
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithDeadline(parent context.Context, d time.Time) (*Context, context.CancelFunc) {
	var c context.CancelFunc
	var ctx = new(Context)
	ctx.Context, c = context.WithDeadline(parent, d)
	return ctx, c
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete:
//
// 	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 		defer cancel()  // releases resources if slowOperation completes before timeout elapses
// 		return slowOperation(ctx)
// 	}
func WithTimeout(parent context.Context, timeout time.Duration) (*Context, context.CancelFunc) {
	var c context.CancelFunc
	var ctx = new(Context)
	ctx.Context, c = context.WithTimeout(parent, timeout)
	return ctx, c
}

// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithCancel(parent context.Context) (*Context, context.CancelFunc) {
	var c context.CancelFunc
	var ctx = new(Context)
	ctx.Context, c = context.WithCancel(parent)
	return ctx, c
}
