// composing interfaces together

package main

import (
	"bytes"
	"fmt"
	"io"
)

type writer interface {
	write([]byte) (int, error)
}

type closer interface {
	close() error
}

type writerCloser interface {
	// embedding interfaces
	writer
	closer
}

/*
since this struct has both write() and close() methods,
it implements both the writer and closer interfaces,
therefore it implements the composed interface writerCloser
*/
type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// this function prints sth to stdout in 8-byte chunks
func (bwc *bufferedWriterCloser) write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		// Write will panic with ErrTooLarge
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// this flushes anything left in the buffer that was not handled by write (for demonstrative purposes)
func (bwc *bufferedWriterCloser) close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// this is a constructor function that returns a pointer to a newly-created bufferedWriterCloser
func newBufferedWriterCloser() *bufferedWriterCloser {
	return &bufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	var wc writerCloser = newBufferedWriterCloser() // note: its underlying type is a bufferWriterCloser
	wc.write([]byte("Hello youtube listeners, this is a test"))
	wc.close()

	// type conversions with interfaces
	bwc := wc.(*bufferedWriterCloser) // can access the buffer directly
	fmt.Println(bwc)

	/*
		cannot convert to an interface that does not implement the same methods
		this will fail:
		bwc := wc.(io.Reader)
		because the writerCloser interface does not wrap the Read method
	*/

	// error handling with comma-ok syntax:
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion to reader failed")
	}
}
