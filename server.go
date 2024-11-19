package maruchi

import "net/http"

type Middleware = func(c Context, next func(c Context))
type Handler = func(c Context)

type Server struct {
	prefix      string
	server      *http.ServeMux
	middlewares []Middleware
}

func NewServer() Server {
	return Server{
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

func (s *Server) handleRequest(index int, handler Handler, c Context) {
	if index >= len(s.middlewares) {
		handler(c)
		return
	}

	s.middlewares[index](c, func(c Context) {
		s.handleRequest(index+1, handler, c)
	})
}

func (s *Server) Route(method string, pattern string, handler Handler) *Server {
	s.server.HandleFunc(method+" "+s.prefix+pattern, func(w http.ResponseWriter, r *http.Request) {
		s.handleRequest(0, handler, &ContextBase{
			ResponseWriter: w,
			Request:        r,
		})
	})

	return s
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
