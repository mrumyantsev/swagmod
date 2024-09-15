package swagmod

import (
	"errors"
	"io"

	"github.com/mrumyantsev/swagmod/internal/config"
	"github.com/mrumyantsev/swagmod/internal/constants"
	filereader "github.com/mrumyantsev/swagmod/internal/file-reader"
	filescanner "github.com/mrumyantsev/swagmod/internal/file-scanner"
	"github.com/mrumyantsev/swagmod/internal/output"
	"github.com/mrumyantsev/swagmod/internal/parser"
	"github.com/mrumyantsev/swagmod/pkg/errlib"
)

type FileScanner interface {
	ScanFiles(dirPath string) ([]constants.DocFile, error)
}

type FileReader interface {
	Reader(filePath string) (io.Reader, error)
}

type Parser interface {
	ParseFile(r io.Reader) ([]byte, error)
}

type App struct {
	config  *config.Config
	scanner FileScanner
	reader  FileReader
	parser  Parser
}

func New() (*App, error) {
	cfg := config.New()

	if err := cfg.Init(); err != nil {
		return nil, err
	}

	return &App{
		config:  cfg,
		scanner: filescanner.New(),
		reader:  filereader.New(),
		parser:  parser.New(),
	}, nil
}

func (a *App) Run() error {
	output.Info("Swagger docs modifying begin...")

	docs, err := a.scanner.ScanFiles(a.config.DocsDir)
	if err != nil {
		return errlib.Wrap(err, "could not scan '"+a.config.DocsDir+"' directory")
	}

	var swagDoc constants.DocFile

	for _, doc := range docs {
		if doc.Type != constants.Json {
			continue
		}

		swagDoc = doc
	}

	if swagDoc.Path == "" {
		return errors.New("no Swagger docs found")
	}

	fileReader, err := a.reader.Reader(swagDoc.Path)
	if err != nil {
		return errlib.Wrap(err, "could not read the file")
	}

	if _, err = a.parser.ParseFile(fileReader); err != nil {
		return errlib.Wrap(err, "could not parse the file")
	}

	return nil
}
