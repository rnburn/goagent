package mock

import (
	"context"
	"fmt"
	"sync"

	"github.com/hypertrace/goagent/sdk"
)

var _ sdk.Span = &Span{}

type Span struct {
	Name       string
	Attributes map[string]interface{}
	Options    sdk.SpanOptions
	err        error
	Noop       bool
	mux        *sync.Mutex
}

func NewSpan() *Span {
	return &Span{mux: &sync.Mutex{}}
}

func (s *Span) SetAttribute(key string, value interface{}) {
	s.mux.Lock() // avoids race conditions
	defer s.mux.Unlock()

	if s.Attributes == nil {
		s.Attributes = make(map[string]interface{})
	}
	s.Attributes[key] = value
}

func (s *Span) ReadAttribute(key string) interface{} {
	s.mux.Lock() // avoids race conditions
	defer s.mux.Unlock()

	val, ok := s.Attributes[key]
	if ok {
		delete(s.Attributes, key)
		return val
	}

	return nil
}

func (s *Span) RemainingAttributes() int {
	return len(s.Attributes)
}

func (s *Span) SetError(err error) {
	s.err = err
}

func (s *Span) IsNoop() bool {
	return s.Noop
}

type spanKey string

func SpanFromContext(ctx context.Context) sdk.Span {
	return ctx.Value(spanKey("span")).(*Span)
}

func StartSpan(ctx context.Context, name string, opts *sdk.SpanOptions) (context.Context, sdk.Span, func()) {
	fmt.Println(opts)
	s := &Span{Name: name, Options: *opts}
	return ContextWithSpan(ctx, s), s, func() {}
}

func ContextWithSpan(ctx context.Context, s sdk.Span) context.Context {
	return context.WithValue(ctx, spanKey("span"), s)
}
