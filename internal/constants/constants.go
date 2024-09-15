package constants

const (
	AppName = "swagmod"
)

type DocType int

const (
	Undefined DocType = iota
	Golang
	Json
	Yaml
)

type DocFile struct {
	Path string
	Type DocType
}

type DocBlock struct {
	Lines       []string
	IsResponses bool
}

type DocBlocks struct {
	Blocks []DocBlock
}
