package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hasitpbhatt/epublib-go/book"
)

func main() {
	input := flag.String("input", "", "input name of the file")

	flag.Parse()

	if input == nil || *input == "" {
		log.Fatal("input file not specified")
	}

	_, err := os.Stat(*input)
	if err != nil {
		log.Fatal("The input file does not exist")
	}

	book, err := book.NewBookFromFile(*input)

	if err != nil {
		log.Fatal(err)
	}

	filesInBook := book.GetPaths()

	for _, path := range filesInBook {
		fmt.Println(path)
	}

	// TODO: Implement operations that could be done on the book
}
