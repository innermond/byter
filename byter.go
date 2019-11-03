package byter

import (
	"io"
	"sync"
)

type Buffer struct {
	io.Writer
	*sync.Mutex
	written int
}

func New(w io.Writer) *Buffer {
	return &Buffer{w, &sync.Mutex{}, 0}
}

func (b *Buffer) WriteString(s string) (n int, err error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	n, err = io.WriteString(b.Writer, s)
	// even in case of an error some bytes,
	// maybe none,
	// has the chance to be written
	b.written += n

	if err != nil {
		return
	}

	return
}

func (b *Buffer) Len() int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	return b.written
}
