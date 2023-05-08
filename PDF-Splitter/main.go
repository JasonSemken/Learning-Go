package main

import (
	"fmt"
	"os"
)

func main() {
	openFile()
}

// Open select file
func openFile() {
	_, err := os.Open("./files/test_file.pdf")
	if err != nil {
		fmt.Println("Error!! ", err)
	}
}

// Read directory and list out file names
func openDir() {
	files, err := os.ReadDir("./files")
	if err != nil {
		fmt.Println("Error!! ", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
