package maruchi

import (
	"context"
	"net/http"
)

type ReqContext interface {
	Context() context.Context
	Req() *http.Request
	Writer() http.ResponseWriter
}

type ReqContextBase struct {
	W http.ResponseWriter
	R *http.Request
}

func (self ReqContextBase) Req() *http.Request {
	return self.R
}

func (self ReqContextBase) Writer() http.ResponseWriter {
	return self.W
}

func (self ReqContextBase) Context() context.Context {
	return self.R.Context()
}
