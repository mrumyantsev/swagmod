package config

import (
	"fmt"
	"strconv"
	"strings"
)

type multiFlags []string

func (m *multiFlags) String() string {
	return "not implemented"
}

func (m *multiFlags) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func parseMappingVals(mapVals []string) (Mappings, error) {
	mappings := make(Mappings, len(mapVals))
	var exit int

	for _, val := range mapVals {
		if exit = parseMappingVal(val, mappings); exit != 0 {
			return nil, fmt.Errorf("invalid mapping '%s'", val)
		}
	}

	return mappings, nil
}

func parseMappingVal(val string, dest Mappings) int {
	if !strings.Contains(val, ":") {
		return 1
	}

	colIdx := strings.Index(val, ":")

	leftSide := val[:colIdx]

	endCode, err := strconv.Atoi(leftSide)
	if err != nil {
		return 1
	}

	rightSide := val[colIdx+1:] // a code range

	codeRange, exit := parseCodeRange(rightSide)
	if exit != 0 {
		return 1
	}

	dest[endCode] = codeRange

	return 0
}

func parseCodeRange(codeRange string) (CodeRange, int) {
	var res CodeRange

	if !strings.Contains(codeRange, "-") {
		code, err := strconv.Atoi(codeRange)
		if err != nil {
			return res, 1
		}

		res.First, res.Last = code, code

		return res, 0
	}

	hypIdx := strings.Index(codeRange, "-")

	firstCode := codeRange[:hypIdx]

	var err error

	res.First, err = strconv.Atoi(firstCode)
	if err != nil {
		return res, 1
	}

	lastCode := codeRange[hypIdx+1:]

	res.Last, err = strconv.Atoi(lastCode)
	if err != nil {
		return res, 1
	}

	return res, 0
}
