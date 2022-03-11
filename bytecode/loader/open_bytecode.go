/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import "os"

/* Open byte code file */
func (loader *Loader) OpenByteCodeFile() error {
	file, err := os.Open(loader.ByteCodeFileName)
	if err != nil {
		return err
	}
	loader.file = file
	return nil
}
