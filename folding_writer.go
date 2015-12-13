package icalparser

import (
	"bytes"
	"unicode/utf8"
)

type FoldingWriter struct {
	b           *bytes.Buffer
	maxLineSize int
	lineSize    int
}

func NewFoldingWriter(maxLineSize int) *FoldingWriter {
	return &FoldingWriter{
		b:           &bytes.Buffer{},
		maxLineSize: maxLineSize,
	}
}

func (f *FoldingWriter) Write(b []byte) (int, error) {
	size := 0
	for len(b) > 0 {
		remain := f.maxLineSize - f.lineSize - 1
		if len(b) > remain {
			for !utf8.Valid(b[:remain]) {
				remain--
			}
			n, _ := f.b.Write(b[:remain])
			size += n
			n, _ = f.b.WriteString("\r\n ")
			size += n
			f.lineSize = 0
			b = b[remain:]
		} else {
			n, _ := f.b.Write(b)
			b = b[n:]
			f.lineSize += n
			size += n
		}
	}

	return size, nil
}

func (f *FoldingWriter) String() string {
	return f.b.String()
}
