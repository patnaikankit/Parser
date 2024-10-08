package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a pdf file path as an arguement")
		os.Exit(1)
	}
}
