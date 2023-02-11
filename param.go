package dune

import (
	"context"
	"net/http"
)

type Param struct {
	key   string
	value string
}

type Params struct {
	i      int
	max    int
	keys   *[]string
	values []string
}

func newParams(cap int) *Params {
	return &Params{
		i:      0,
		max:    cap,
		keys:   nil,
		values: make([]string, 0, cap),
	}
}

func (p *Params) setKeys(keys *[]string) {
	p.keys = keys
	p.values = p.values[:len(*keys)]
}

func (p *Params) appendValue(value string) {
	if p.i >= p.max {
		return
	}
	p.values[p.i] = value
	p.i++
}

func (p *Params) reset() {
	p.i = 0
	p.keys = nil
	p.values = p.values[:0]
}

func (p *Params) Get(k string) string {
	if p.keys != nil {
		for i, key := range *p.keys {
			if key == k {
				return p.values[i]
			}
		}
	}
	return ""
}

func (p *Params) Copy() *Params {
	cp := new(Params)
	*cp = *p
	return cp
}

func (p *Params) ForEach(f func(k, v string) bool) {
	if p.keys != nil {
		for i := len(*p.keys) - 1; i >= 0; i-- {
			f((*p.keys)[i], p.values[i])
		}
	}
}

var paramKey = &struct{}{}

func withParamsCtx(ctx context.Context, ps *Params) context.Context {
	return context.WithValue(ctx, paramKey, ps)
}

func hasParamsCtx(ctx context.Context) bool {
	return ctx.Value(paramKey) != nil
}

func paramsFromCtx(ctx context.Context) (*Params, bool) {
	ps, ok := ctx.Value(paramKey).(*Params)
	return ps, ok
}

// emptyParams is a Params object with 0 capacity, therefore its basically immutable and concurrent safe.
var emptyParams = newParams(0)

// Vars returns the http.Request's Params from which route variables can be retrieved.
func Vars(r *http.Request) *Params {
	if ps, ok := paramsFromCtx(r.Context()); ok {
		return ps
	}

	return emptyParams
}
