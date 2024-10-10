package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/patnaikankit/Parser/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a pdf file path as an arguement")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	absPath, _ := filepath.Abs(pdfPath)

	if err := pkg.CheckFileExists(absPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text, err := pkg.ExtractText(absPath)
	if err != nil {
		fmt.Println("Error Extracting Text")
		os.Exit(1)
	}

	parsedDoc := pkg.ParseText(text)
	pkg.RenderParsedText(parsedDoc)
}
