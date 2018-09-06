# Mserv

## Example

```go
func main() {
    s := mserv.New(
        // pprof
        mserv.NewHTTPServer(time.Second, &http.Server{
            Addr:    ":8081",
            Handler: http.DefaultServeMux,
        }),
        // prometheus
        mserv.NewHTTPServer(time.Second, &http.Server{
            Addr:    ":8082",
            Handler: promhttp.Handler(),
        }),
        // gin
        mserv.NewHTTPServer(5*time.Second, &http.Server{
            Addr:         ":8083",
            Handler:      ginApp(),
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 10 * time.Second,
        }),
        // echo
        mserv.NewHTTPServer(5*time.Second, &http.Server{
            Addr:         ":8084",
            Handler:      echoApp(),
            TLSConfig:    &tls.Config{ /**todo**/ },
            ReadTimeout:  5 * time.Second,
            WriteTimeout: 5 * time.Second,
        }),
        // grpc
        mserv.NewGRPCServer(":8085", grpcServer()),
    )

    // start all servers
    s.Start()
    // wait stop signal
    <-grace.ShutdownContext(context.Background()).Done()
    // graceful stop each server
    s.Stop()
}
```

full example [here](/examples/main.go)
