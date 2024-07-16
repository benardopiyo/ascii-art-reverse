package utils

import (
	"testing"
)

func TestDivideTemplatesIntoSegments(t *testing.T) {
	// Define the input template lines
	templateLines := []string{
		"Header", // This will be skipped in segmenting
		"Line 1", "Line 2", "Line 3", "Line 4", "Line 5", "Line 6", "Line 7", "Line 8", "Line 9",
		"Line 10", "Line 11", "Line 12", "Line 13", "Line 14", "Line 15", "Line 16", "Line 17", "Line 18",
	}

	// Define the expected segments
	expectedSegments := [][]string{
		{"Line 1", "Line 2", "Line 3", "Line 4", "Line 5", "Line 6", "Line 7", "Line 8", "Line 9"},
		{"Line 10", "Line 11", "Line 12", "Line 13", "Line 14", "Line 15", "Line 16", "Line 17", "Line 18"},
	}

	segments := DivideTemplatesIntoSegments(templateLines)

	if len(segments) != len(expectedSegments) {
		t.Fatalf("Expected %d segments, got %d", len(expectedSegments), len(segments))
	}

	// Check the content of each segment
	for i, segment := range segments {
		for j, line := range segment {
			if line != expectedSegments[i][j] {
				t.Errorf("Expected segment %d line %d to be %q, got %q", i, j, expectedSegments[i][j], line)
			}
		}
	}
}

func TestIsSegmentPresent(t *testing.T) {
	segmentLines := []string{"Line 1", "Line 2", "Line 3"}                                        // Define the segment lines to check
	targetLines := []string{"Line 1 More Text", "Line 2 More Text", "Line 3 More Text", "Line 4"} // Define the target lines

	// Check if the segment is present at the beginning of the target lines
	if !IsSegmentPresent(segmentLines, targetLines) {
		t.Error("Expected segment to be present")
	}

	segmentLines = []string{"Line 1", "Line 2", "Line X"} // Modify the segment lines to a non-matching case
	// Check if the modified segment is not present
	if IsSegmentPresent(segmentLines, targetLines) {
		t.Error("Expected segment to not be present")
	}
}

func TestRemoveSegment(t *testing.T) {
	segmentLength := 6                                                                            // Define the length of the segment to be removed
	targetLines := []string{"Line 1 More Text", "Line 2 More Text", "Line 3 More Text", "Line 4"} // Define the target lines
	expectedLines := []string{" More Text", " More Text", " More Text", "Line 4"}                 // Define the expected lines after the segment is removed

	modifiedLines := RemoveSegment(segmentLength, targetLines)

	if len(modifiedLines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(modifiedLines))
	}

	for i, line := range modifiedLines {
		if line != expectedLines[i] {
			t.Errorf("Expected line %d to be %q, got %q", i, expectedLines[i], line)
		}
	}
}
