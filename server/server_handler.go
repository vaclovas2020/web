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
