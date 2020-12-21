// Package that contains all the default routines
// and type to request/response the rest api
package utils

import (
	"encoding/json"
	"net/http"

	"./types"
)

// DefaultMarshal is responsible to format all the
// response object
func DefaultsMarshal(w http.ResponseWriter, code string, message string) ([]byte, error) {
	resp := &types.ResponseError{
		Code:    code,
		Message: message,
	}

	arr, err := json.Marshal(resp)
	Check(w, code, err)

	return arr, nil
}
