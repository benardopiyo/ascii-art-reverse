package utils

import (
	"os"
	"testing"
)

func TestFileReader(t *testing.T) {
	// Setup: Create a temporary file with test content
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the file afterwards

	// Write some content to the temporary file
	content := "Line 1\nLine 2\nLine 3\n"
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Call the function being tested
	lines, err := FileReader(tmpfile.Name())
	if err != nil {
		t.Fatalf("FileReader returned an error: %v", err)
	}

	// Verify the results
	expectedLines := []string{"Line 1", "Line 2", "Line 3", ""}
	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Expected line %d to be %q, got %q", i, expectedLines[i], line)
		}
	}
}
