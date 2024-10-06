package middlewares

import (
	"log/slog"
	"net/http"
)

type recoverMiddlewareWriter struct {
  http.ResponseWriter
  headerWritten bool
}

func (w *recoverMiddlewareWriter) WriteHeader(code int) {
  if w.headerWritten {
    slog.Warn("Header already written")
    return
  }

  w.ResponseWriter.WriteHeader(code)
  w.headerWritten = true
}

func RecoverMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    rmw := &recoverMiddlewareWriter{w, false}

    defer func() {
      if re := recover(); re != nil {
        slog.Warn("Recover from panic:", re)
        http.Error(w,
          http.StatusText(http.StatusInternalServerError),
          http.StatusInternalServerError,
        )
      }
    }()

    next.ServeHTTP(rmw, r)
  })
}
