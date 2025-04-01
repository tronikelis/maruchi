package maruchi

import "net/http"

type (
	Middleware = func(r ReqContext, next func(r ReqContext))
	Handler    = func(r ReqContext)
)

type Server struct {
	prefix      string
	server      *http.ServeMux
	middlewares []Middleware
}

func NewServer() *Server {
	return &Server{
		server:      http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

// copies the current middlewares registered
func (s *Server) Group(pattern string) *Server {
	middlewares := make([]Middleware, len(s.middlewares))
	copy(middlewares, s.middlewares)

	return &Server{
		server:      s.server,
		prefix:      s.prefix + pattern,
		middlewares: middlewares,
	}
}

func (s *Server) handleRequest(index int, handler Handler, r ReqContext) {
	if index >= len(s.middlewares) {
		handler(r)
		return
	}

	s.middlewares[index](r, func(r ReqContext) {
		s.handleRequest(index+1, handler, r)
	})
}

func (s *Server) Route(method string, pattern string, handler Handler) *Server {
	s.server.HandleFunc(method+" "+s.prefix+pattern, func(w http.ResponseWriter, r *http.Request) {
		s.handleRequest(0, handler, ReqContextBase{
			writer:  w,
			request: r,
		})
	})

	return s
}

// connect http.Handler with maruchi,
// this is like .ServeMux().Handle, but with middlewares
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.server.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		s.handleRequest(
			0,
			func(r ReqContext) {
				handler.ServeHTTP(r.Writer(), r.Req())
			},
			ReqContextBase{writer: w, request: r},
		)
	})
}

func (s *Server) GET(pattern string, handler Handler) *Server {
	return s.Route("GET", pattern, handler)
}

func (s *Server) POST(pattern string, handler Handler) *Server {
	return s.Route("POST", pattern, handler)
}

func (s *Server) PUT(pattern string, handler Handler) *Server {
	return s.Route("PUT", pattern, handler)
}

func (s *Server) DELETE(pattern string, handler Handler) *Server {
	return s.Route("DELETE", pattern, handler)
}

func (s *Server) Middleware(fn Middleware) *Server {
	s.middlewares = append(s.middlewares, fn)
	return s
}

// returns the prefix string for all requests
func (s *Server) Prefix() string {
	return s.prefix
}

// usually you will call http.ListenAndServer(addr, server.ServeMux())
// at the end of your configuration
func (s *Server) ServeMux() *http.ServeMux {
	return s.server
}
