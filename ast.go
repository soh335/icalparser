package icalparser

import (
	"bytes"
	"io"
)

type Object struct {
	HeaderLine     *ContentLine
	PropertiyLines []*ContentLine
	Components     []*Component
	FooterLine     *ContentLine
}

func (o *Object) String() string {
	var b bytes.Buffer
	b.WriteString(o.HeaderLine.String())
	b.WriteString("\r\n")
	for _, line := range o.PropertiyLines {
		b.WriteString(line.String())
		b.WriteString("\r\n")
	}
	for _, line := range o.Components {
		b.WriteString(line.String())
	}
	b.WriteString(o.FooterLine.String())
	b.WriteString("\r\n")
	return b.String()
}

type Component struct {
	HeaderLine     *ContentLine
	PropertiyLines []*ContentLine
	FooterLine     *ContentLine
}

func (c *Component) String() string {
	var b bytes.Buffer
	b.WriteString(c.HeaderLine.String())
	b.WriteString("\r\n")
	for _, line := range c.PropertiyLines {
		b.WriteString(line.String())
		b.WriteString("\r\n")
	}
	b.WriteString(c.FooterLine.String())
	b.WriteString("\r\n")
	return b.String()
}

type ContentLine struct {
	Name  *Ident
	Param []*Param
	Value *Ident
}

func (c *ContentLine) String() string {
	f := NewFoldingWriter(75)
	io.WriteString(f, c.Name.String())

	for i, param := range c.Param {
		if i < len(c.Param) {
			io.WriteString(f, ";")
		}
		io.WriteString(f, param.String())
	}

	io.WriteString(f, ":")
	io.WriteString(f, c.Value.String())
	return f.String()
}

type Param struct {
	ParamName   *Ident
	ParamValues []*Ident
}

func (p *Param) String() string {
	var b bytes.Buffer
	b.WriteString(p.ParamName.String())
	for i, paramValue := range p.ParamValues {
		b.WriteString(paramValue.String())
		if i < len(p.ParamValues)-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

type Ident struct {
	C     string
	Token token
}

func (i *Ident) String() string {
	var b bytes.Buffer
	switch i.Token {
	case tokenQuotedString:
		b.WriteString(`"`)
		b.WriteString(i.C)
		b.WriteString(`"`)
	default:
		b.WriteString(i.C)
	}

	return b.String()
}
