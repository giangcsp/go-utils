package httpUtils

import (
  "io"
	"net/http"

	"github.com/bytedance/sonic"
)

// Unmarshal request body in json format to dst
func FromJson(r *http.Request, dst interface{}) error {
  body, err := io.ReadAll(r.Body)
  if err != nil {
    return err
  }

  if err := sonic.Unmarshal(body, dst); err != nil {
    return err
  }

  return nil
}
