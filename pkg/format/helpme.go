// Package format provides utilities for formatting help messages for command-line applications.
// It includes functionality to format flags and their descriptions in a uniform manner,
// ensuring that the output is neatly aligned and easy to read.
package format

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

// templateHelpMe is a Go template that iterates over a collection of items,
// formatting each item with a flag and its corresponding description. The flag
// is left-aligned with a specified width, and the description is indented by
// the same width. This template is useful for generating formatted help or
// usage messages where flags and their descriptions need to be aligned neatly.
var templateHelpMe = `
{{- range .}}
  {{ printf "%-*s" .Width .Flag }}   {{ .Description | indent .Width}}
{{- end }}
`

// HelpMe represents a structure that holds information about a command-line flag.
// It includes the flag name, a description of the flag, and the width for formatting purposes.
type HelpMe struct {
	Flag        string
	Description string
	Width       int
}

// PrintHelpMe formats and prints a slice of HelpMe structs using a predefined template.
// It first calculates the width for each HelpMe item, then sets up a template with custom functions
// for formatting. The template is parsed and executed, rendering the formatted output to os.Stdout.
//
// Parameters:
//
//	helpme []HelpMe - A slice of HelpMe structs to be formatted and printed.
func PrintHelpMe(helpme []HelpMe) {
	width := parseWidth(helpme)
	for i := range helpme {
		helpme[i].Width = width
	}

	// funcMap is a map of functions that can be used in the template.
	// In this case, we are using the printf function from the fmt package to format the output.
	funcMap := template.FuncMap{
		"printf": fmt.Sprintf,
		"indent": indent,
	}

	// tmpl is a template object that contains the help output template.
	// The template is a string that contains the formatting for the help output.
	tmpl, err := template.New("helpme").Funcs(funcMap).Parse(templateHelpMe)
	if err != nil {
		panic(err)
	}

	// The Execute function is used to render the template and print the formatted output.
	// The first argument is the output destination (os.Stdout) and the second argument is the data model (helpme).
	err = tmpl.Execute(os.Stdout, helpme)
	if err != nil {
		panic(err)
	}

	fmt.Println("teste")
}

// parseWidth takes a slice of HelpMe structs and returns the length of the longest
// flag string within the slice. It iterates through each HelpMe struct, trims any
// leading or trailing whitespace from the flag string, and updates the maximum
// width if the current flag string is longer than the previously recorded maximum width.
func parseWidth(flags []HelpMe) int {
	maxWidth := 0

	for _, flag := range flags {
		// Trim leading and trailing whitespace from the flag.Flag string
		flag.Flag = strings.TrimSpace(flag.Flag)

		// Check if the length of the flag.Flag string is greater than maxWidth
		// If it is, update maxWidth with the new length
		if len(flag.Flag) > maxWidth {
			maxWidth = len(flag.Flag)
		}
	}

	return maxWidth
}

// indent formats a multi-line string by adding indentation to each line except the first one.
// The indentation is calculated based on the provided width and a fixed flag width of 4 spaces.
// It trims any leading or trailing whitespace from the input description before processing.
// If the description contains only one line, it returns the original description without modification.
//
// Parameters:
//   - width: The number of spaces to add as indentation.
//   - description: The multi-line string to be indented.
//
// Returns:
//
//	A new string with the specified indentation applied to each line except the first one.
func indent(width int, description string) string {
	trimSpaceDescription := strings.TrimSpace(description)
	flagWidth := 4

	lines := strings.Split(trimSpaceDescription, "\n")
	if len(lines) <= 1 {
		return description
	}

	// The for loop iterates over the lines of the description and adds indentation to each line.
	// The indentation is calculated based on the width of the flag and a fixed value (4).
	for idx := 1; idx < len(lines); idx++ {
		lines[idx] = strings.Repeat(" ", width+flagWidth) + lines[idx]
	}

	return strings.Join(lines, "\n")
}
