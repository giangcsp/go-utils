package middlewares

import (
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
		latency := end.Sub(start)
		if latency > time.Minute {
			latency.Truncate(time.Second)
		}

		slog.LogAttrs(r.Context(), slog.LevelInfo,
			"Logger",
			slog.Int("statusCode", lmw.statusCode),
			slog.Duration("latency", latency),
			slog.String("address", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)
	})
}
