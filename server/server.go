/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package server

import (
	"fmt"
	"log"
	"net/http"

	"webimizer.dev/web/base"
)

type Server struct {
	Host        string
	Port        int
	ServerClass *base.Class
	RouterClass *base.Class
}

func (sr Server) Start() {
	log.Printf("\033[32m[weblang]\033[0m Server starting listen on %v:%v...", sr.Host, sr.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", sr.Host, sr.Port), nil))
}
