package controller

import (
	"net/http"
)

type Handler func(rw http.ResponseWriter, r *http.Request)

func (handler Handler) Invoke(args map[string]interface{}) error {
	// TODO: Invoke handler
	return nil
}
