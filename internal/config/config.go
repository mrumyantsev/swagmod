package config

import (
	"flag"

	"github.com/mrumyantsev/swagmod/internal/constants"
)

type CodeRange struct {
	First int
	Last  int
}

type Mappings map[int]CodeRange

type Config struct {
	DocsDir  string
	Mappings Mappings
}

func New() *Config {
	return &Config{}
}

func (c *Config) Init() error {
	mappingVals := new(multiFlags)

	flag.Var(mappingVals, "m", constants.MFlagUsage)

	flag.Parse()

	mappings, err := parseMappingVals(*mappingVals)
	if err != nil {
		return err
	}

	c.Mappings = mappings

	c.DocsDir = "./docs"

	return nil
}
