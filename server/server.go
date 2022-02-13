/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package server

import (
	"fmt"
	"log"
	"net/http"

	"webimizer.dev/web/base"
)

type Server struct {
	Host         string
	Port         int
	ServerObject *base.Object
	RouterObject *base.Object
}

func (sr Server) Start() error {
	log.Printf("\033[32m[weblang]\033[0m Server starting listen on %v:%v...", sr.Host, sr.Port)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", sr.Host, sr.Port), nil)
}
