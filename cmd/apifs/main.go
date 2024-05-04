package main

import (
	"log"
	"os"

	"github.com/howardplus/apifs/pkg/apifs"
)

func main() {
	args := os.Args

	fpath := "/data"
	if len(args) > 1 {
		fpath = args[1]
	}

	p, err := apifs.NewProcessor(fpath, "/api/v1/", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to create apifs processor, err %v", err)
	}

	if err := p.Run(); err != nil {
		log.Fatalf("failed to start server, err %v", err)
	}

	log.Printf("program exiting...")
}
