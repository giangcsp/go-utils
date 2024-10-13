package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type CorsConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	MaxAge         int
}

var DefaultCorsConfig CorsConfig = CorsConfig{
	[]string{"*"},
	[]string{"*"},
	[]string{"*"},
	86400,
}

func MakeCorsMiddleware(cfg CorsConfig) Middleware {
	allowedOriginsString := strings.Join(cfg.AllowedOrigins, ", ")
	allowedMethodsString := strings.Join(cfg.AllowedMethods, ", ")
	allowedHeadersString := strings.Join(cfg.AllowedHeaders, ", ")
	maxAgeString := fmt.Sprintf("%d", cfg.MaxAge)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOriginsString)
			w.Header().Set("Access-Control-Allow-Methods", allowedMethodsString)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeadersString)
			w.Header().Set("Access-Control-Max-Age", maxAgeString)
			
      // Preflight request
			if _, ok := r.Header["Origin"]; r.Method == "OPTIONS" && ok {
				w.WriteHeader(204)
				return
			}
      
      next.ServeHTTP(w, r)
		})
	}
}
