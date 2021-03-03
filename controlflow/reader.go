package controlflow

import (
	"errors"
	"fmt"
	"io"
)

type BadReader struct {
	err error
}
type SimpleReader struct {
	count int
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	println(sr.count)
	if sr.count > 3 {
		return 0, io.EOF
	}
	sr.count++
	return sr.count, nil
}

func ReadSomething() error {
	var r io.Reader = BadReader{err: errors.New("Does not exist")}

	if _, err := r.Read([]byte("test something")); err != nil {
		fmt.Printf("An error occurred: %s", err)
		return err
	}
	return nil
}
