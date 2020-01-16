package services

import (
	"os"
	"path/filepath"

	"github.com/zWaR/gostego/pkg/gostego/interfaces"
)

type file struct{}

// PathExists checks if a file exists
func (file *file) FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// BaseNoExt returns the basename of the filepath without its extension
func (file *file) BaseNoExt(path string) string {
	basename := filepath.Base(path)
	extension := filepath.Ext(basename)
	filenameLength := len(basename) - len(extension)
	return basename[:filenameLength]
}

// NewFileService returns a new FileService
func NewFileService() interfaces.FileService {
	var fileInstance = new(file)
	var fileService interfaces.FileService = fileInstance
	return fileService
}
