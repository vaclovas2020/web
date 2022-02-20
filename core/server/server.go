/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Server implemantation for use in Weblang VM runtime environment  */
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"webimizer.dev/web/base"
	"webimizer.dev/web/core/cms"
	"webimizer.dev/webimizer"
)

/* Main server struct */
type Server struct {
	cms             cms.CMS      // Content Managment System
	Host            string       // Server hostname
	Port            int          // Server port
	StaticFilesPath string       // Static files path (optional)
	ServerObject    *base.Object // Pointer to server object in VM environment
	RouterObject    *base.Object //  Pointer to router object in VM environment
}

/* Start server process */
func (sr Server) Start() error {
	webimizer.DefaultHTTPHeaders = [][]string{
		{"x-content-type-options", "nosniff"},
		{"x-frame-options", "SAMEORIGIN"},
		{"x-xss-protection", "1; mode=block"},
	} // define default headers
	sr.initStaticFilesHandler()
	if err := sr.cms.ServeStaticFiles(); err != nil {
		return err
	}
	log.Printf("\033[32m[weblang]\033[0m Server starting listen on %v:%v...", sr.Host, sr.Port)
	sr.generateCmsAdminUrl()
	return http.ListenAndServe(fmt.Sprintf("%v:%v", sr.Host, sr.Port), nil)
}

/* Generate CMS admin url */
func (sr *Server) generateCmsAdminUrl() {
	adminUrl := "/admin" + uuid.New().String() + "/"
	sr.cms = cms.CMS{AdminUrl: adminUrl}
	log.Printf("\033[32m[weblang]\033[0m Your admin CMS url is http://%v:%v%v", sr.Host, sr.Port, adminUrl)
}
