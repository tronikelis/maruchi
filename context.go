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
	writer  http.ResponseWriter
	request *http.Request
}

func (self ReqContextBase) Writer() http.ResponseWriter {
	return self.writer
}

func (self ReqContextBase) Req() *http.Request {
	return self.request
}

func (self ReqContextBase) Context() context.Context {
	return self.request.Context()
}
