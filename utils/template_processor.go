package utils

// Divide the given slice of strings (TemplateLines) into segments of 9 lines each
// Segmenting starts from line2 (index 1) and continues until the end of slice
func DivideTemplatesIntoSegments(TemplateLines []string) [][]string {
	var segments [][]string

	for startIndex := 1; startIndex < len(TemplateLines); startIndex += 9 {
		endIndex := startIndex + 9 // Calculate the end index for the current segment
		if endIndex > len(TemplateLines) {
			endIndex = len(TemplateLines)
		}

		segment := TemplateLines[startIndex:endIndex]
		segments = append(segments, segment)
	}
	return segments
}

// Checks if the segment represented by segmentLines
// is present at the beginning of the targetLines.
func IsSegmentPresent(segmentLines, targetLines []string) bool {

	if len(segmentLines[0]) > len(targetLines[0]) { // sLines cannot be a prefix of tLines if L1 of sL is longer than  L1 of tLines
		return false
	}

	// Exclude the last line, help not to run out of bounds when comparing with segmentLines.
	for i, line := range targetLines[:len(targetLines)-1] {
		if segmentLines[i] != line[:len(segmentLines[i])] {
			return false
		}
	}
	return true
}

// Removes the matched segment from the target lines by
// trimming the matched segment's length from each line.
func RemoveSegment(segmentLength int, targetLines []string) []string {
	for i, line := range targetLines[:len(targetLines)-1] {
		targetLines[i] = line[segmentLength:]
	}
	return targetLines
}
