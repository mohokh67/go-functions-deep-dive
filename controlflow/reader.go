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

func (sr *SimpleReader) Close() error {
	println("closing reader...")
	return nil
}

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	if sr.count == 3 {
		panic("Something very very bad happended")
	}
	if sr.count > 5 {
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

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		// Work like First In Last Out , FILO - So we can have multiple defer func
		// Always run after returning value in a function
		// e.g. could be file close, db connection close etc.
		_ = r.Close()
		if p := recover(); p != nil {
			println(p)
			err = errors.New("a panic happenned but recovered. It's ok for now")
		}
	}()

	defer func() {
		// not a good practice, we don't need multiple panic.
		println("defer func before loop")
	}()

	for {
		value, readerErr := r.Read([]byte("I am the content of a file"))
		if readerErr == io.EOF {
			println("Finish reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println(value)
	}
	defer func() {
		// not a good practice, we don't need multiple panic.
		println("defer func after loop")
	}()

	return nil

}
