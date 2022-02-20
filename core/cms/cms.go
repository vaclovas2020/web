/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang Content Managment System */
package cms

import (
	"embed"
	"errors"
	"net/http"

	"webimizer.dev/webimizer"
)

/* Content managment system struct */
type CMS struct {
	AdminUrl string
}

// content is our static web server content.
//go:embed static/*
var content embed.FS

func (cms *CMS) ServeStaticFiles() error {
	if cms.AdminUrl == "" {
		return errors.New("admin url cannot be empty string")
	}
	http.Handle(cms.AdminUrl,
		http.StripPrefix(cms.AdminUrl,
			webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(content)).ServeHTTP(rw, r) })))
	return nil
}
