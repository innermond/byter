package byter

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestBuffer_Memory(t *testing.T) {
	store := &bytes.Buffer{}

	buf := New(store)

	test(buf, t)
}

func TestBuffer_File(t *testing.T) {

	store, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(store.Name())

	buf := New(store)

	test(buf, t)
}

func test(buf *Buffer, t *testing.T) {
	if buf.Len() != 0 {
		t.Errorf("expected a newly created buffer to have 0 length, got %d\n", buf.Len())
	}

	// first write
	n, err := buf.WriteString(" ")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	got := buf.Len()
	if n != int(got) {
		t.Errorf("expected %d got %d\n", n, got)
	}

	nn := n
	for i := 0; i < 10; i++ {
		n, err = buf.WriteString(" ")
		if err != nil {
			t.Errorf("%v\n", err)
		}

		nn += n

		got = buf.Len()
		if nn != int(got) {
			t.Errorf("expected %d got %d\n", n, got)
		}
	}
}
