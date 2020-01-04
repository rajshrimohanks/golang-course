package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Get the file path
	filePath := os.Args[1]

	// Open the file
	file, fOpenError := os.Open(filePath)
	if fOpenError != nil {
		fmt.Println("Error:", fOpenError)
		os.Exit(1)
	}

	// Write to terminal
	io.Copy(os.Stdout, file)
}
