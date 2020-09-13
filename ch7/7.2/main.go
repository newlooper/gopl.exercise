package main

import "io"

type CounterWriter struct {
	writer  io.Writer
	counter int64
}

func (cw *CounterWriter) Write(p []byte) (int, error) {
	cw.counter += int64(len(p))
	return cw.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{w, 0}
	return &cw, &cw.counter
}
