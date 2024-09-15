package filewriter

import (
	"os"

	"github.com/mrumyantsev/swagmod/pkg/errlib"
	fsops "github.com/mrumyantsev/swagmod/pkg/fs-ops"
)

const (
	eol      = "\n"
	filePerm = 0644
)

type FileWriter struct {
}

func New() *FileWriter {
	return &FileWriter{}
}

func (w *FileWriter) AppendToFile(filePath, entry string) error {
	err := fsops.MakeFileDirectories(filePath)
	if err != nil {
		return errlib.Wrap(err, "could not create directories for file")
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
	if err != nil {
		return errlib.Wrap(err, "could not open file for line appending")
	}
	defer file.Close()

	if _, err := file.WriteString(entry + eol); err != nil {
		return errlib.Wrap(err, "could not append a line into the file")
	}

	return nil
}
