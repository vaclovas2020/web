/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package web

import (
	"errors"
)

/* Install CMS to database and set userName and password */
func (vm *VM) InstallCmsUser(userName, password string) error {
	var connString string
	for _, conf := range vm.config.Config {
		if conf.Name == "monog.connectionString" {
			connString = conf.Value
			break
		}
	}
	if connString == "" {
		return errors.New("mongodb connection string is empty")
	}
	return nil
}
