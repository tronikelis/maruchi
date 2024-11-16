package maruchi

import (
	"context"
	"net/http"
)

type Context interface {
	http.ResponseWriter

	Context() context.Context
	SetContextVal(key any, val any)
	ContextVal(key any) any
	Req() *http.Request
}

type ContextBase struct {
	http.ResponseWriter
	Request *http.Request
}

func (c *ContextBase) Req() *http.Request {
	return c.Request
}

func (c *ContextBase) SetContextVal(key any, val any) {
	ctx := context.WithValue(c.Request.Context(), key, val)
	c.Request = c.Request.WithContext(ctx)
}

func (c *ContextBase) ContextVal(key any) any {
	return c.Request.Context().Value(key)
}

func (c *ContextBase) Context() context.Context {
	return c.Request.Context()
}
