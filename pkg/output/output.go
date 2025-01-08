// Package output provides utilities for capturing and redirecting standard output.
// It allows capturing the output of functions that write to os.Stdout and returning
// the captured output as a string. This can be useful for testing and logging purposes.
package output

import (
	"bytes"
	"io"
	"os"
)

// Output represents a structure that holds file descriptors for standard output,
// and read/write operations. It contains the following fields:
// - stdout: a file descriptor for standard output.
// - read: a file descriptor for reading operations.
// - write: a file descriptor for writing operations.
type Output struct {
	stdout *os.File
	read   *os.File
	write  *os.File
}

// Capture creates a new Output instance with a pipe for capturing stdout.
// It returns a pointer to the Output struct containing the original stdout,
// and the read and write ends of the pipe.
func Capture() *Output {
	read, write, _ := os.Pipe()

	return &Output{
		stdout: os.Stdout,
		read:   read,
		write:  write,
	}
}

// New captures the output of a function that writes to os.Stdout and returns it as a string.
// It temporarily redirects os.Stdout to a pipe, executes the provided captureOutput function,
// and then restores os.Stdout to its original state. The captured output is read from the pipe
// and returned as a string.
//
// Parameters:
//
//	captureOutput - a function that writes output to os.Stdout.
//
// Returns:
//
//	A string containing the captured output.
func (o *Output) New(captureOutput func()) string {
	var buf bytes.Buffer
	os.Stdout = o.write

	captureOutput()

	o.write.Close()

	os.Stdout = o.stdout

	io.Copy(&buf, o.read)
	return buf.String()
}
