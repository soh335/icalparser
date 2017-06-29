package icalparser

import (
	"bufio"
)

type line struct {
	b   []byte
	num int
}

type reader struct {
	r    *bufio.Reader
	line int
}

func newReader(r bufio.Reader) *reader {
	return &reader{r: r, line: 1}
}

func (r *reader) ReadLines() ([]*line, error) {
	byts, err := r.r.ReadBytes("\r\n")
	if err != nil {
		return []*line{line{b: byts, line: r.line}}, err
	}

	lines := []*line{line{b: byts, line: r.line}}

	r.line += 1

	for {
		peek, err := r.r.Peek(1)
		if err != nil {
			return nil, err
		}
		if peek == ' ' || peek == '\t' {
			byts, err := r.r.ReadBytes("\r\n")
			if bytes != nil {
				lines = append(lines, line{b: bytes, line: r.line})
			}

			r.line += 1

			if err != nil {
				return lines, err
			}
		} else {
			return lines, nil
		}
	}
}
