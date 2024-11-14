# maruchi

A bare-bones wrapper around `http.ServeMux`

```go
// initialize the server
// it uses `http.ServeMux` under the hood
server := maruchi.NewServer()

// add middleware
server.Middleware(func(c maruchi.Context, next func(c maruchi.Context)) {
    // pre request
    // can pass your own context
    next(c)
    // post request
})

server.GET("/user/{id}", func(c maruchi.Context) {
    c.Writer().Write([]byte("hello world"))
})

// grouping
// middleware gets copied over
group := server.Group("/auth")

group.GET("/login", func(c maruchi.Context) {
    // --
})
```
