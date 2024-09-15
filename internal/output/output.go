package output

import (
	"os"

	"github.com/mrumyantsev/swagmod/internal/constants"
)

const (
	colSpc = ": "
	eol    = "\n"

	fatalExitCode = 1
)

func Info(msg string) {
	os.Stdout.WriteString(msg + eol)
}

func Fatal(err error) {
	os.Stderr.WriteString(constants.AppName + colSpc + err.Error() + eol)

	os.Exit(fatalExitCode)
}
