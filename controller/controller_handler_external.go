package controller

import (
	"errors"
	"fmt"
	"net/http"

	"webimizer.dev/web/base"
	app "webimizer.dev/webimizer"
)

type Handler func(rw http.ResponseWriter, r *http.Request)

/* Invoke Http handler to serve request */
func (handler Handler) Invoke(args map[string]interface{}, funcPtr *base.Function, obj *base.Object) error {
	if v, found := args["route"]; found {
		if method, found := args["method"]; found {
			methodStr := fmt.Sprintf("%v", method)
			http.Handle(fmt.Sprintf("%v", v), app.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
				switch methodStr {
				case http.MethodHead:
					app.Head(rw, r, app.IfHttpMethodHandler(handler))
				case http.MethodGet:
					app.Get(rw, r, app.IfHttpMethodHandler(handler))
				case http.MethodPost:
					app.Post(rw, r, app.IfHttpMethodHandler(handler))
				case http.MethodPut:
					app.Put(rw, r, app.IfHttpMethodHandler(handler))
				case http.MethodDelete:
					app.Delete(rw, r, app.IfHttpMethodHandler(handler))
				case http.MethodOptions:
					app.Options(rw, r, app.IfHttpMethodHandler(handler))
				}
			}))
		} else {
			http.Handle(fmt.Sprintf("%v", v), app.HttpHandler(handler))
		}
	} else {
		return errors.New("controller has no route defined")
	}
	return nil
}
