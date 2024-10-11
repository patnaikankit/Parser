package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/patnaikankit/Parser/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a pdf file path as an argument")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	absPath, _ := filepath.Abs(pdfPath)

	if err := pkg.CheckFileExists(absPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	paragraphs, err := pkg.ExtractText(absPath)
	if err != nil {
		fmt.Printf("Error Extracting Text: %v\n", err)
		os.Exit(1)
	}

	err = pkg.RenderParsedTextToFile(paragraphs, "output.txt")
	if err != nil {
		fmt.Printf("Error Rendering Text: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Text extracted and written to output.txt")
}
