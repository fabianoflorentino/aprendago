// Package output provides utilities for capturing and redirecting standard output.
// It allows capturing the output of functions that write to os.Stdout and returning
// the captured output as a string. This can be useful for testing and logging purposes.
package output

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
)

// Output represents a structure that holds file descriptors for standard output,
// and read/write operations. It contains the following fields:
// - stdout: a file descriptor for standard output.
// - read: a file descriptor for reading operations.
// - write: a file descriptor for writing operations.
type Output struct {
	stdout *os.File
	stderr *os.File
	read   *os.File
	write  *os.File
}

// New creates a new Output instance with a pipe for reading and writing.
// It returns a pointer to the Output struct with stdout set to the standard output,
// and read and write set to the ends of the pipe.
func New() *Output {
	read, write, err := os.Pipe()
	if err != nil {
		logger.Log("Failed to create pipe: %v", err)
	}

	return &Output{
		stdout: os.Stdout,
		stderr: os.Stderr,
		read:   read,
		write:  write,
	}
}

// Capture redirects the standard output to a buffer, executes the provided
// captureOutput function, and then restores the standard output. It returns
// the captured output as a string. This is useful for capturing and testing
// output generated by functions that write to standard output.
func (o *Output) Capture(captureOutput func()) (string, error) {
	var buf bytes.Buffer
	var wg sync.WaitGroup

	// Redireciona Stdout e Stderr
	os.Stdout = o.write
	os.Stderr = o.write
	log.SetOutput(o.write)

	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(&buf, o.read)
	}()

	captureOutput()

	o.write.Close()

	wg.Wait()

	os.Stdout = o.stdout
	os.Stderr = o.stderr
	log.SetOutput(os.Stderr)

	return buf.String(), nil
}
