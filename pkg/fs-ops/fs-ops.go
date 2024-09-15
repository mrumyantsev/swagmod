package fsops

import (
	"os"
	"strings"
)

const (
	dirCreatePerm = 0755
)

// IsDirectory determines if a file represented by path is a directory
// or not.
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// IsDirectoryExist checks for the directory's existence and returns
// true if it exists and false if not.
func IsDirectoryExist(dirPath string) bool {
	_, err := os.Stat(dirPath)

	return os.IsExist(err)
}

// MakeDirectory creates a directory specified by path and returns
// error if it failed.
func MakeDirectory(path string) error {
	if err := os.Mkdir(path, dirCreatePerm); err != nil {
		return err
	}

	return nil
}

func MakeFileDirectories(filePath string) error {
	idx := strings.LastIndex(filePath, "/")

	dirPath := filePath[:idx]

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, dirCreatePerm); err != nil {
			return err
		}
	}

	return nil
}
