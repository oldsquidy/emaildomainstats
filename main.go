package main

import (
	"emaildomainstats/emaildomainstats"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("input file not supplied")
		return
	}

	reader, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}
	defer reader.Close()

	if err := emaildomainstats.ProcessData(reader, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
