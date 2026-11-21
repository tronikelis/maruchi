# maruchi

A bare-bones wrapper around `http.ServeMux`

```go
// initialize the muxer
// it uses `http.ServeMux` under the hood
muxer := maruchi.NewMuxer()

// add middleware
muxer.Middleware(func(r maruchi.ReqContext, next func(c maruchi.ReqContext)) {
    // pre request
    // can pass your own context
    next(c)
    // post request
})

muxer.GET("/user/{id}", func(r maruchi.ReqContext) {
    c.Write([]byte("hello world"))
})

// grouping
// middleware gets copied over
group := muxer.Group("/auth")

group.GET("/login", func(r maruchi.ReqContext) {
    // --
})

// simply
http.ListenAndServe("localhost:3000", muxer.ServeMux())

// or with configuration
server := &http.Server{Handler: muxer.ServeMux()}
```
