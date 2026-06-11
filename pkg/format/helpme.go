// Package format provides utilities for formatting help messages for command-line applications.
package format

// HelpMe represents a structure that holds information about a command-line flag.
type HelpMe struct {
	Flag        string
	Description string
	Width       int
}
