package main

import (
	"golang.org/x/tour/reader"
	//"fmt"
	//"strings"
	//"io"
	)

type MyReader struct{
	r string
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
	return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}
