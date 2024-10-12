package middlewares

import (
  "net/http"
)

type Middleware func(next http.Handler) http.Handler

func WithMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
  h := handler
  for i := len(middlewares) - 1; i >= 0; i-- {
    h = middlewares[i](h)
  }

  return h
}
