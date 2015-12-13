package icalparser

import (
	"bufio"
	"fmt"
	"io"
)

type Parser struct {
	s *scanner
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: &scanner{r: bufio.NewReader(r)}}
}

func (p *Parser) Parse() (*Object, error) {
	object := &Object{
		components:     []*Component{},
		propertiyLines: []*ContentLine{},
	}

	lineNum := 0
	var component *Component

	for {
		line, err := p.parseLine()

		if object.footerLine != nil && err == io.EOF {
			return object, nil
		} else {
			if line == nil {
				return nil, err
			}
			if err != nil {
				return nil, err
			}
		}

		switch lineNum {
		case 0:
			if e, g := "BEGIN:VCALENDAR", line.AsString(); e != g {
				return nil, fmt.Errorf("first line is expected %v but got %v", e, g)
			}
			object.headerLine = line
		default:
			// object properties
			if component == nil {
				if line.name.c == "BEGIN" {
					component = &Component{
						headerLine:     line,
						propertiyLines: []*ContentLine{},
					}
				} else if line.AsString() == "END:VCALENDAR" {
					object.footerLine = line
				} else {
					if len(object.components) == 0 {
						object.propertiyLines = append(object.propertiyLines, line)
					} else {
						return nil, fmt.Errorf("unexpected line: %s", line.AsString())
					}
				}
			} else {
				if line.AsString() == "END:"+component.headerLine.value.c {
					component.footerLine = line
					object.components = append(object.components, component)
					component = nil
				} else {
					component.propertiyLines = append(component.propertiyLines, line)
				}
			}
		}
		lineNum++
	}
}

func (p *Parser) parseLine() (*ContentLine, error) {
	contentLine := &ContentLine{}

	name, token := p.s.scanName()

	if err := p.s.err; err != nil {
		return nil, err
	}

	contentLine.name = &Ident{c: name, token: token}

	paramList := []*Param{}
	for p.s.ch == ';' {
		param, err := p.parseParam()
		if err != nil {
			return nil, err
		}
		paramList = append(paramList, param)
	}
	contentLine.param = paramList

	switch p.s.ch {
	case ':':
		value, token := p.s.scanValue()
		contentLine.value = &Ident{c: value, token: token}
	default:
		return nil, fmt.Errorf(`unexpected error: should got : but %v`, string(p.s.ch))
	}

	switch err := p.s.err; {
	case err == nil:
		if !p.s.isCRLF() {
			return nil, fmt.Errorf(`should \r\n`)
		}
		p.s.read() // \n
		return contentLine, nil
	case err == io.EOF:
		return contentLine, err
	default:
		return nil, err
	}
}

func (p *Parser) parseParam() (*Param, error) {
	param := &Param{}

	paramName, token := p.s.scanParamName()
	if p.s.err != nil {
		return nil, p.s.err
	}

	param.paramName = &Ident{c: paramName, token: token}

	switch p.s.ch {
	case '=':
		paramValues := []*Ident{}
	OUTER:
		for {
			paramValue, token := p.s.scanParamValue()
			if err := p.s.err; err != nil {
				return nil, err
			}
			paramValues = append(paramValues, &Ident{c: paramValue, token: token})
			switch p.s.ch {
			case ',':
			default:
				break OUTER
			}
		}
		param.paramValues = paramValues
		return param, p.s.err
	default:
		return nil, fmt.Errorf(`unexpected error: should got = but %s`, string(p.s.ch))
	}
}
