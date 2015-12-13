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

func (o *Object) AsString() string {
	var b bytes.Buffer
	b.WriteString(o.headerLine.AsString())
	b.WriteString("\r\n")
	for _, line := range o.propertiyLines {
		b.WriteString(line.AsString())
		b.WriteString("\r\n")
	}
	for _, line := range o.components {
		b.WriteString(line.AsString())
	}
	b.WriteString(o.footerLine.AsString())
	b.WriteString("\r\n")
	return b.String()
}

type Component struct {
	headerLine     *ContentLine
	propertiyLines []*ContentLine
	footerLine     *ContentLine
}

func (c *Component) AsString() string {
	var b bytes.Buffer
	b.WriteString(c.headerLine.AsString())
	b.WriteString("\r\n")
	for _, line := range c.propertiyLines {
		b.WriteString(line.AsString())
		b.WriteString("\r\n")
	}
	b.WriteString(c.footerLine.AsString())
	b.WriteString("\r\n")
	return b.String()
}

type ContentLine struct {
	name  *Ident
	param []*Param
	value *Ident
}

func (c *ContentLine) AsString() string {
	f := NewFoldingWriter(75)
	io.WriteString(f, c.name.AsString())

	for i, param := range c.param {
		if i < len(c.param) {
			io.WriteString(f, ";")
		}
		io.WriteString(f, param.AsString())
	}

	io.WriteString(f, ":")
	io.WriteString(f, c.value.AsString())
	return f.String()
}

type Param struct {
	paramName   *Ident
	paramValues []*Ident
}

func (p *Param) AsString() string {
	var b bytes.Buffer
	b.WriteString(p.paramName.AsString())
	for i, paramValue := range p.paramValues {
		b.WriteString(paramValue.AsString())
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

func (i *Ident) AsString() string {
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
