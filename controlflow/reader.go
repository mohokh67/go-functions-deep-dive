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

func ReadFullFile() error {
	var r io.Reader = &SimpleReader{}
	defer func() {
		// Work like First In Last Out , FILO - So we can have multiple defer func
		// Always run after returning value in a function
		// e.g. could be file close, db connection close etc.
		println("I am a defer function, ")
	}()

	for {
		value, readerErr := r.Read([]byte("I am the content of a file"))
		if readerErr == io.EOF {
			println("Finish reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			return readerErr
		}
		println(value)
	}

	return nil

}
