package chain_interceptors

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

const key = "foo"

func getStringFromContext(ctx context.Context, key string) string {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}

// Returns the request (as a string), concatenated to v, concatenated to the
// value associated with key in ctx (as a string).
func makeTestHandler(v string) Handler {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		value := getStringFromContext(ctx, key)
		return req.(string) + v + value, nil
	}
}

func makeTestInterceptor(prependToRsp, appendToCtx string) Interceptor {
	return func(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
		value := getStringFromContext(ctx, key)
		rsp, err := handler(context.WithValue(ctx, key, value+appendToCtx), req)
		return prependToRsp + rsp.(string), err
	}
}

type chainer func(fns ...Interceptor) Interceptor

var testChainers = []struct {
	name string
	fn   chainer
}{
	{"chain1", chain1},
	{"chain2", chain2},
	{"chain3", chain3},
	{"chain4", chain4},
}

var testCases = []struct {
	interceptors []Interceptor
	handler      Handler
	req, resp    string
}{
	{
		[]Interceptor{
			makeTestInterceptor("a", "b"),
		},
		makeTestHandler("hdlr_"),
		"_REQ",
		"a_REQhdlr_b",
	},
	{
		[]Interceptor{
			makeTestInterceptor("a", "b"),
			makeTestInterceptor("c", "d"),
		},
		makeTestHandler("hdlr_"),
		"_REQ",
		"ac_REQhdlr_bd",
	},
	{
		[]Interceptor{
			makeTestInterceptor("a", "b"),
			makeTestInterceptor("c", "d"),
			makeTestInterceptor("e", "f"),
		},
		makeTestHandler("hdlr_"),
		"_REQ",
		"ace_REQhdlr_bdf",
	},
	{
		[]Interceptor{
			makeTestInterceptor("a", "b"),
			makeTestInterceptor("c", "d"),
			makeTestInterceptor("e", "f"),
			makeTestInterceptor("g", "h"),
		},
		makeTestHandler("hdlr_"),
		"_REQ",
		"aceg_REQhdlr_bdfh",
	},
}

func TestChainers(t *testing.T) {
	for _, chainer := range testChainers {
		for i, tt := range testCases {
			t.Run(fmt.Sprintf("%s_%d", chainer.name, i), func(t *testing.T) {
				chained := chainer.fn(tt.interceptors...)
				rsp, err := chained(context.Background(), tt.req, tt.handler)
				if err != nil {
					t.Fatal("chained handler returned an error:", err)
				}
				s, ok := rsp.(string)
				if !ok {
					t.Fatalf("response is not a string: %+v", rsp)
				}
				if s != tt.resp {
					t.Fatalf("response is %q, want %q", s, tt.resp)
				}
			})
		}
	}
}

func TestChainersConcurrent(t *testing.T) {
	const N = 100
	for _, chainer := range testChainers {
		for i, tt := range testCases {
			t.Run(fmt.Sprintf("concurrent_%s_%d", chainer.name, i), func(t *testing.T) {
				chained := chainer.fn(tt.interceptors...)
				var wg sync.WaitGroup
				for i := 0; i < N; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for i := 0; i < N; i++ {
							rsp, err := chained(context.Background(), tt.req, tt.handler)
							if err != nil {
								panic(fmt.Sprintf("chained handler returned an error: %+v", err))
							}
							s, ok := rsp.(string)
							if !ok {
								panic(fmt.Sprintf("response is not a string: %+v", rsp))
							}
							if s != tt.resp {
								panic(fmt.Sprintf("response is %q, want %q", s, tt.resp))
							}
						}
					}()
				}
				wg.Wait()
			})
		}
	}
}

func h0(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}
func i0(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}
func i1(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}
func i2(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}
func i3(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}
func i4(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}
func i5(ctx context.Context, req interface{}, handler Handler) (interface{}, error) {
	return handler(ctx, req)
}

var gResp interface{} // to break loop optimization

func BenchmarkChainers(b *testing.B) {
	tests := [][]Interceptor{
		{i0},
		{i0, i1},
		{i0, i1, i2},
		{i0, i1, i2, i3},
		{i0, i1, i2, i3, i4},
		{i0, i1, i2, i3, i4, i5},
	}
	for _, interceptors := range tests {
		for _, chainer := range testChainers {
			b.Run(fmt.Sprintf("%s_interceptors_%d", chainer.name, len(interceptors)), func(b *testing.B) {
				chained := chainer.fn(interceptors...)
				for i := 0; i < b.N; i++ {
					resp, err := chained(context.Background(), "REQ", h0)
					if err != nil {
						b.Fatalf("%s(%d) returned an error: %v", chainer.name, len(interceptors), err.Error())
					}
					gResp = resp
				}
			})
		}
	}
}
