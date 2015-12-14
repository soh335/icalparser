package icalparser

import (
	"bytes"
	"io"
)

type Object struct {
	headerLine     *ContentLine
	propertiyLines []*ContentLine
	components     []*Component
	footerLine     *ContentLine
}

func (o *Object) String() string {
	var b bytes.Buffer
	b.WriteString(o.headerLine.String())
	b.WriteString("\r\n")
	for _, line := range o.propertiyLines {
		b.WriteString(line.String())
		b.WriteString("\r\n")
	}
	for _, line := range o.components {
		b.WriteString(line.String())
	}
	b.WriteString(o.footerLine.String())
	b.WriteString("\r\n")
	return b.String()
}

type Component struct {
	headerLine     *ContentLine
	propertiyLines []*ContentLine
	footerLine     *ContentLine
}

func (c *Component) String() string {
	var b bytes.Buffer
	b.WriteString(c.headerLine.String())
	b.WriteString("\r\n")
	for _, line := range c.propertiyLines {
		b.WriteString(line.String())
		b.WriteString("\r\n")
	}
	b.WriteString(c.footerLine.String())
	b.WriteString("\r\n")
	return b.String()
}

type ContentLine struct {
	name  *Ident
	param []*Param
	value *Ident
}

func (c *ContentLine) String() string {
	f := NewFoldingWriter(75)
	io.WriteString(f, c.name.String())

	for i, param := range c.param {
		if i < len(c.param) {
			io.WriteString(f, ";")
		}
		io.WriteString(f, param.String())
	}

	io.WriteString(f, ":")
	io.WriteString(f, c.value.String())
	return f.String()
}

type Param struct {
	paramName   *Ident
	paramValues []*Ident
}

func (p *Param) String() string {
	var b bytes.Buffer
	b.WriteString(p.paramName.String())
	for i, paramValue := range p.paramValues {
		b.WriteString(paramValue.String())
		if i < len(p.paramValues)-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

type Ident struct {
	c     string
	token token
}

func (i *Ident) String() string {
	var b bytes.Buffer
	switch i.token {
	case tokenQuotedString:
		b.WriteString(`"`)
		b.WriteString(i.c)
		b.WriteString(`"`)
	default:
		b.WriteString(i.c)
	}

	return b.String()
}
