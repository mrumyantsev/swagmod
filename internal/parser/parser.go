package parser

import (
	"bufio"
	"io"
	"strings"

	"github.com/mrumyantsev/swagmod/internal/constants"
	linesplitter "github.com/mrumyantsev/swagmod/internal/parser/line-splitter"
)

type Parser struct {
	splitter *linesplitter.LineSplitter
}

func New() *Parser {
	return &Parser{
		splitter: linesplitter.New(),
	}
}

func (p *Parser) ParseFile(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(p.splitter.Split)

	var line string
	var fours int
	var lineCloseBracket string
	var isInResponses bool

	var blocks constants.DocBlocks
	var block constants.DocBlock

	for scanner.Scan() {
		line = scanner.Text()

		if strings.Contains(line, "responses") {
			fours = strings.Count(line, "    ")
			lineCloseBracket = strings.Repeat("    ", fours) + "}"

			isInResponses = true
			block.IsResponses = true
		} else if line == lineCloseBracket {
			isInResponses = false
			block.IsResponses = false
		}

		block.Lines = append(block.Lines, line)
	}

	// remove
	_ = isInResponses
	_ = blocks

	return []byte{}, nil
}
