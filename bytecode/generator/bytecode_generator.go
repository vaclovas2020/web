/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"webimizer.dev/web/bytecode"
)

/* Weblang bytecode class file generator for use in Weblang VM */
type ByteCodeGenerator struct {
	Data     bytecode.ByteCode // data struct
	FileName string            // Weblang bytecode full file name
}
