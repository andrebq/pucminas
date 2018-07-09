package extract

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

type (
	// Pipe extracts a list of "number number" with
	// volts measurements and mAmps measurements and outputs
	// a CSV line by attaching the timestamp to the input pair.
	//
	// Empty lines are removed
	Pipe struct {
		in  *bufio.Scanner
		out io.Writer
	}
)

// NewPipe returns a new pipe
func NewPipe(out io.Writer, in io.Reader) *Pipe {
	return &Pipe{
		in:  bufio.NewScanner(in),
		out: out,
	}
}

// WriteHeader writes the CSV header to output
func (p *Pipe) WriteHeader() error {
	_, err := fmt.Fprintf(p.out, "ts,mAmps\n")
	return err
}

// Copy reads one line from input and writes it to the output
func (p *Pipe) Copy() error {
	for p.in.Scan() {
		line := bytes.Trim(p.in.Bytes(), "\n\r")
		if len(line) == 0 {
			continue
		}
		var mamps int
		_, err := fmt.Fscanf(bytes.NewBuffer(line), "%d", &mamps)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(p.out, "%v,%v\n", time.Now().Format(time.RFC3339),
			mamps)
		if err != nil {
			return err
		}
	}
	return p.in.Err()
}
