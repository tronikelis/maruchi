package maruchi

import (
	"context"
	"net/http"
)

type ReqContext interface {
	http.ResponseWriter

	Context() context.Context
	SetContextVal(key any, val any)
	ContextVal(key any) any
	Req() *http.Request
}

type ReqContextBase struct {
	http.ResponseWriter
	Request *http.Request
}

func (r *ReqContextBase) Req() *http.Request {
	return r.Request
}

func (r *ReqContextBase) SetContextVal(key any, val any) {
	ctx := context.WithValue(r.Request.Context(), key, val)
	r.Request = r.Request.WithContext(ctx)
}

func (r *ReqContextBase) ContextVal(key any) any {
	return r.Request.Context().Value(key)
}

func (r *ReqContextBase) Context() context.Context {
	return r.Request.Context()
}
