package chain_interceptors

import (
	"context"
)

// These are roughly equivalent to gRPC handlers and interceptors.
type Handler func(ctx context.Context, req interface{}) (interface{}, error)
type Interceptor func(ctx context.Context, req interface{}, handler Handler) (interface{}, error)

func chain1(fns ...Interceptor) Interceptor {
	return func(ctx context.Context, req interface{}, handler Handler) (resp interface{}, err error) {
		var (
			h Handler
			i = 0 // current interceptor is fns[i]
		)
		h = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			i++
			if i < len(fns) {
				return fns[i](ctx, req, h)
			}
			return handler(ctx, req)
		}
		return fns[0](ctx, req, h)
	}
}

func chain2(fns ...Interceptor) Interceptor {
	n := len(fns)
	curr := fns[n-1] // last interceptor is innermost
	for i := n - 2; i >= 0; i-- {
		var (
			prev = curr
			fn   = fns[i]
		)
		curr = func(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
			h := func(ctx context.Context, req interface{}) (interface{}, error) { return prev(ctx, req, handler) }
			return fn(ctx, req, h)
		}
	}
	return curr
}

func chain3(fns ...Interceptor) Interceptor {
	combine := func(i1, i2 Interceptor) Interceptor {
		return func(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
			return i1(ctx, req, func(ctx context.Context, req interface{}) (interface{}, error) {
				return i2(ctx, req, handler)
			})
		}
	}
	acc := fns[0]
	for _, fn := range fns[1:] {
		acc = combine(acc, fn)
	}
	return acc
}

func chain4(fns ...Interceptor) Interceptor {
	if len(fns) == 1 {
		return fns[0]
	}
	if len(fns) == 2 {
		f0, f1 := fns[0], fns[1]
		return func(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
			return f0(ctx, req, func(ctx context.Context, req interface{}) (interface{}, error) {
				return f1(ctx, req, handler)
			})
		}
	}
	if len(fns) == 3 {
		f0, f1, f2 := fns[0], fns[1], fns[2]
		return func(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
			return f0(ctx, req, func(ctx context.Context, req interface{}) (interface{}, error) {
				return f1(ctx, req, func(ctx context.Context, req interface{}) (interface{}, error) {
					return f2(ctx, req, handler)
				})
			})
		}
	}
	return chain1(fns...)
}
