package middlewares

import (
  "net/http"
)

type Middleware func(next http.Handler) http.Handler

func WithMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
  h := handler
  for _, m := range middlewares {
    h = m(h)
  }

  return h
}
