# maruchi

A bare-bones wrapper around `http.ServeMux`

```go
// initialize the server
// it uses `http.ServeMux` under the hood
server := maruchi.NewServer()

// add middleware
server.Middleware(func(r maruchi.ReqContext, next func(c maruchi.ReqContext)) {
    // pre request
    // can pass your own context
    next(c)
    // post request
})

server.GET("/user/{id}", func(r maruchi.ReqContext) {
    c.Write([]byte("hello world"))
})

// grouping
// middleware gets copied over
group := server.Group("/auth")

group.GET("/login", func(r maruchi.ReqContext) {
    // --
})

http.ListenAndServe("localhost:3000", server.ServeMux())
```
