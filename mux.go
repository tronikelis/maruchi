package maruchi

import "net/http"

type (
	Middleware = func(r ReqContext, next func(r ReqContext))
	Handler    = func(r ReqContext)
)

type Muxer struct {
	prefix      string
	server      *http.ServeMux
	middlewares []Middleware
}

func NewMuxer() *Muxer {
	return &Muxer{
		server:      http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

// copies the current middlewares registered
func (self *Muxer) Group(pattern string) *Muxer {
	middlewares := make([]Middleware, len(self.middlewares))
	copy(middlewares, self.middlewares)

	return &Muxer{
		server:      self.server,
		prefix:      self.prefix + pattern,
		middlewares: middlewares,
	}
}

func (self *Muxer) handleRequest(index int, handler Handler, r ReqContext) {
	if index >= len(self.middlewares) {
		handler(r)
		return
	}

	self.middlewares[index](r, func(r ReqContext) {
		self.handleRequest(index+1, handler, r)
	})
}

func (self *Muxer) Route(method string, pattern string, handler Handler) *Muxer {
	self.server.HandleFunc(method+" "+self.prefix+pattern, func(w http.ResponseWriter, r *http.Request) {
		self.handleRequest(0, handler, ReqContextBase{
			W: w,
			R: r,
		})
	})

	return self
}

// connect http.Handler with maruchi,
// this is like .ServeMux().Handle, but with middlewares
func (self *Muxer) Handle(pattern string, handler http.Handler) {
	self.server.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		self.handleRequest(
			0,
			func(r ReqContext) {
				handler.ServeHTTP(r.Writer(), r.Req())
			},
			ReqContextBase{W: w, R: r},
		)
	})
}

func (self *Muxer) GET(pattern string, handler Handler) *Muxer {
	return self.Route("GET", pattern, handler)
}

func (self *Muxer) POST(pattern string, handler Handler) *Muxer {
	return self.Route("POST", pattern, handler)
}

func (self *Muxer) PUT(pattern string, handler Handler) *Muxer {
	return self.Route("PUT", pattern, handler)
}

func (self *Muxer) DELETE(pattern string, handler Handler) *Muxer {
	return self.Route("DELETE", pattern, handler)
}

func (self *Muxer) Middleware(fn Middleware) *Muxer {
	self.middlewares = append(self.middlewares, fn)
	return self
}

// returns the prefix string for all requests
func (self *Muxer) Prefix() string {
	return self.prefix
}

func (self *Muxer) ServeMux() *http.ServeMux {
	return self.server
}
