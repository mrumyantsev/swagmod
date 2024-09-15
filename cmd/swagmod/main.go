package main

import (
	"github.com/mrumyantsev/swagmod/internal/app/swagmod"
	"github.com/mrumyantsev/swagmod/internal/output"
)

func main() {
	app, err := swagmod.New()
	if err != nil {
		output.Fatal(err)
	}

	if err = app.Run(); err != nil {
		output.Fatal(err)
	}
}
