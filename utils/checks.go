package utils

import (
	"errors"
	"net/http"
	"strings"
)

//Check generic error response
func Check(w http.ResponseWriter, code string, e error) (http.ResponseWriter, error) {
	if e != nil {
		arr, _ := DefaultsMarshal(w, code, e.Error())
		w.Write(arr)
		w.WriteHeader(400)
		return w, e
	}
	return w, nil
}

//CheckContentType return wrong content-type error
func CheckContentType(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		return Check(w, "http-1", errors.New("Invalid Content-type"))
	}

	return w, nil
}

//CheckMethod return wrong http method error
func CheckMethod(w http.ResponseWriter, r *http.Request, method string) (http.ResponseWriter, error) {
	if r.Method != method {
		return Check(w, "http-2", errors.New("Invalid http method"))
	}

	return w, nil
}
