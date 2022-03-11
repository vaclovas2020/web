package loader

import (
	"fmt"
)

func (loader *Loader) readData(size int64) ([]byte, error) {
	if loader.filePos+size > (*loader.fileStat).Size() {
		return nil, fmt.Errorf("EOF reached when try to read from '%v' file", loader.ByteCodeFileName)
	}
	data := make([]byte, size)
	count, err := loader.file.ReadAt(data, loader.filePos)
	if err != nil {
		return nil, err
	}
	loader.filePos += int64(count)
	return data, nil
}
