package maruchi

import (
	"context"
	"net/http"
)

type Context interface {
	Context() context.Context
	SetContextVal(key any, val any)
	ContextVal(key any) any

	Req() *http.Request
	Writer() http.ResponseWriter
}

type ContextBase struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *ContextBase) Req() *http.Request {
	return c.Request
}

func (c *ContextBase) Writer() http.ResponseWriter {
	return c.ResponseWriter
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
