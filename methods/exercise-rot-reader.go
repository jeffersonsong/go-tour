package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (n int, err error) {
	n, err = reader.r.Read(p)

	if err == nil {
		for i := 0; i < n; i++ {
			p[i] = rot13(p[i])
		}
	}

	return n, err
}

func rot13(c byte) byte {
	switch {
	case c >= 'A' && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case c >= 'a' && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
