package main

import (
	"os"

	"github.com/howardplus/apifs/pkg/apifs"
)

func main() {
	p, err := apifs.NewProcessor("/data", "/api/v1/", 8080)
	if err != nil {
		os.Exit(-1)
	}

	p.Run()
}
