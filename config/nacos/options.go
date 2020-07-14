package nacos

import (
	"context"

	"github.com/micro/go-micro/v2/config/source"
)

type endpointKey struct{}
type namespaceIDKey struct{}
type accessKey struct{}
type secretKey struct{}
type dataIDKey struct{}
type groupKey struct{}

// WithEndpoint ...
func WithEndpoint(e string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, endpointKey{}, e)
	}
}

// WithNamespace ...
func WithNamespace(n string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, namespaceIDKey{}, n)
	}
}

// WithAccountKey ...
func WithAccountKey(a string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, accessKey{}, a)
	}
}

// WithSecretKey ...
func WithSecretKey(s string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, secretKey{}, s)
	}
}

// WithDataIDKey ...
func WithDataIDKey(s string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, dataIDKey{}, s)
	}
}

// WithGroupKey ...
func WithGroupKey(s string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, groupKey{}, s)
	}
}
