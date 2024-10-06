package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type loggerMiddlewareWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggerMiddlewareWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.statusCode = code
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lmw := &loggerMiddlewareWriter{w, 200}

		next.ServeHTTP(lmw, r)

		end := time.Now()
		go func() {
			latency := end.Sub(start)
			if latency > time.Minute {
				latency.Truncate(time.Second)
			}

			slog.Info(fmt.Sprintf("%v |%3d| %13v | %15s | %s %s\n",
				start.Format("2006/01/02 - 15:04:05"),
				lmw.statusCode,
				latency,
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
			))
		}()
	})
}
