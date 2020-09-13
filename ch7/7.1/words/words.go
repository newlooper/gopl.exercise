package words

import (
	"bufio"
	"bytes"
	"fmt"
)

// word counter
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	// get runes
	s := bufio.NewScanner(bytes.NewBuffer(p))

	// word split
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}

	return len(p), s.Err()
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d word(s)", *c)
}

// line counter
type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	for s.Scan() {
		*l++
	}
	return len(p), s.Err()
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("%d line(s)", *l)
}
