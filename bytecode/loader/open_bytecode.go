/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"
	"os"
)

/* Open bytecode file */
func (loader *Loader) openByteCodeFile() error {
	file, err := os.Open(loader.ByteCodeFileName)
	if err != nil {
		return err
	}
	loader.file = file
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("openByteCodeFile: %v", err.Error())
	}
	loader.fileStat = &stat
	return nil
}
