/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package server

import (
	"net/http"

	"webimizer.dev/webimizer"
)

func (sr Server) initStaticFilesHandler() {
	if sr.StaticFilesPath != "" {
		http.Handle("/", webimizer.NewFileServerHandler(sr.StaticFilesPath))
	}
}
