package utils

import (
	"log"
	"os"
	"strings"
)

// Read the content of the files and returns a slice of strings
// where each string represents a line in the file

func FileReader(filename string) ([]string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return nil, err
	}

	lines := strings.Split(string(fileContent), "\n")
	return lines, err
}
