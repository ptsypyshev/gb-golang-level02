package filesCreator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateFile(dirPath, fileName, content string) error {
	fullPath := filepath.Join(dirPath, fileName)
	newFile, err := os.Create(fullPath)
	if err != nil {
		return CreateFileError{
			path:       fullPath,
			innerError: err.Error(),
			errorTime:  time.Now(),
		}
	}
	defer newFile.Close() // Only affect NIX systems, at Windows 10 works fine without Close()

	if _, err := newFile.WriteString(content); err != nil {
		return fmt.Errorf("cannot write to file %s", fullPath)
	}

	return nil
}
