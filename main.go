package main

import (
	"fmt"
	"os"
	"strings"

	"reverse/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 || !strings.HasPrefix(args[0], "--reverse=") {
		fmt.Println("Usage: go run . [OPTION]\n\ngo run . --reverse=<FileName>")
		return
	}

	standardTemplate, err := utils.FileReader("standard.txt")
	if err != nil {
		fmt.Printf("Error reading standard template %v\n", err)
		return
	}

	segmentedStdTemplate := utils.DivideTemplatesIntoSegments(standardTemplate)

	targetFile := args[0][10:] // Get the name of the target file

	targetTemplate, err := utils.FileReader(targetFile)
	if err != nil {
		fmt.Printf("Error reading target template file '%s' : %v\n :", targetFile, err)
		return
	}

	var reversedText strings.Builder // Variable to store the reversed text.

	// Process the target template until it is empty.
	for len(targetTemplate) > 0 && len(targetTemplate[0]) > 0 {
		for i, segment := range segmentedStdTemplate {
			if utils.IsSegmentPresent(segment, targetTemplate) {
				reversedText.WriteRune(rune(i + 32))                                  // Append the corresponding char to the reversed text.
				targetTemplate = utils.RemoveSegment(len(segment[0]), targetTemplate) // Remove the matched segment from the target template.
			}
		}
	}

	fmt.Println(reversedText.String())
}
