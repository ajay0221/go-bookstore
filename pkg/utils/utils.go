package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err != nil {
		return err
	} else {
		return json.Unmarshal(body, x)
	}
}
