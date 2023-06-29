package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 'A'
	}

	return len(buf), nil
}

func main() {
	reader.Validate(MyReader{})
}
