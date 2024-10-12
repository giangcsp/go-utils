package main

import (
  "log"
  "log/slog"
  "net/http"
  middlewares "github.com/giangcsp/go-utils/pkg/http/middlewares"
)

func main() {
  server := http.NewServeMux()

  server.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
  })

  server.HandleFunc("GET /divideByZero", func (w http.ResponseWriter, r *http.Request) {
    b := 1
    b--
    _ = 1 / b
    w.Write([]byte("Something went wrong"))
  })

  slog.Info("Serving on port 6900")
  log.Fatal(
    http.ListenAndServe(
      ":6900", 
      middlewares.WithMiddlewares(
        server, 
        middlewares.LoggerMiddleware, 
        middlewares.RecoverMiddleware,
      ),
    ),
  )
}
