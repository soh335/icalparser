package icalparser

import (
	"io"
)

type Printer struct {
	o *Object
}

func NewPrinter(o *Object) *Printer {
	return &Printer{o}
}

func (p *Printer) WriteTo(w io.Writer) (int, error) {
	return io.WriteString(w, p.o.AsString())
}
