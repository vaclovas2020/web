/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import "os"

/* Open bytecode file */
func (loader *Loader) openByteCodeFile() error {
	file, err := os.Open(loader.ByteCodeFileName)
	if err != nil {
		return err
	}
	loader.file = file
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	loader.fileStat = &stat
	return nil
}
