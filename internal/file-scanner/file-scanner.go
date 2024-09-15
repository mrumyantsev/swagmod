package filescanner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mrumyantsev/swagmod/internal/constants"
	"github.com/mrumyantsev/swagmod/pkg/errlib"
)

type FileScanner struct {
	filePaths []constants.DocFile
}

func New() *FileScanner {
	return &FileScanner{}
}

func (s *FileScanner) ScanFiles(dirPath string) ([]constants.DocFile, error) {
	if err := filepath.Walk(dirPath, s.walker); err != nil {
		return nil, errlib.Wrap(err, "could not walk the directory")
	}

	return s.filePaths, nil
}

func (s *FileScanner) walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	doc := constants.DocFile{
		Path: path,
		Type: docType(path),
	}

	s.filePaths = append(s.filePaths, doc)

	return nil
}

func docType(path string) constants.DocType {
	dotIdx := strings.LastIndexByte(path, '.')

	if dotIdx < 0 {
		return constants.Undefined
	}

	ext := strings.ToLower(path[dotIdx+1:])

	switch ext {
	case "go":
		return constants.Golang
	case "json":
		return constants.Json
	case "yml":
		fallthrough
	case "yaml":
		return constants.Yaml
	}

	return constants.Undefined
}
