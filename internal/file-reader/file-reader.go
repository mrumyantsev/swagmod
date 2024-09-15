package filereader

import (
	"bytes"
	"io"
	"os"

	"github.com/mrumyantsev/swagmod/pkg/errlib"
)

type FileReader struct {
}

func New() *FileReader {
	return &FileReader{}
}

func (r *FileReader) Reader(filePath string) (io.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errlib.Wrap(err, "could not open a file")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errlib.Wrap(err, "could not read all data from the file")
	}

	return bytes.NewReader(data), nil
}
