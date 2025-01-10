// Package trim provides functionality to handle and manipulate strings by trimming leading and trailing white spaces.
// It defines a Space struct to encapsulate the string and provides methods to initialize and trim the string.
package trim

import "strings"

// Space represents a structure that holds a string with trimmed spaces.
type Space struct{}

// Trim takes a string s and returns a Space struct containing the string.
// It is used to initialize a Space struct with the provided string.
func New() *Space {
	return &Space{}
}

// New trims leading and trailing white spaces from the string stored in the Space struct.
// It returns the trimmed string.
func (s Space) String(txt string) string {
	return strings.TrimSpace(txt)
}
