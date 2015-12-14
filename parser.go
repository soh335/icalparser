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
		Components:     []*Component{},
		PropertiyLines: []*ContentLine{},
	}

	lineNum := 0
	var component *Component

	for {
		line, err := p.parseLine()

		if object.FooterLine != nil && err == io.EOF {
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
			if e, g := "BEGIN:VCALENDAR", line.String(); e != g {
				return nil, fmt.Errorf("first line is expected %v but got %v", e, g)
			}
			object.HeaderLine = line
		default:
			// object properties
			if component == nil {
				if line.Name.C == "BEGIN" {
					component = &Component{
						HeaderLine:     line,
						PropertiyLines: []*ContentLine{},
					}
				} else if line.String() == "END:VCALENDAR" {
					object.FooterLine = line
				} else {
					if len(object.Components) == 0 {
						object.PropertiyLines = append(object.PropertiyLines, line)
					} else {
						return nil, fmt.Errorf("unexpected line: %s", line.String())
					}
				}
			} else {
				if line.String() == "END:"+component.HeaderLine.Value.C {
					component.FooterLine = line
					object.Components = append(object.Components, component)
					component = nil
				} else {
					component.PropertiyLines = append(component.PropertiyLines, line)
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

	contentLine.Name = &Ident{C: name, Token: token}

	paramList := []*Param{}
	for p.s.ch == ';' {
		param, err := p.parseParam()
		if err != nil {
			return nil, err
		}
		paramList = append(paramList, param)
	}
	contentLine.Param = paramList

	switch p.s.ch {
	case ':':
		value, token := p.s.scanValue()
		contentLine.Value = &Ident{C: value, Token: token}
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

	param.ParamName = &Ident{C: paramName, Token: token}

	switch p.s.ch {
	case '=':
		paramValues := []*Ident{}
	OUTER:
		for {
			paramValue, token := p.s.scanParamValue()
			if err := p.s.err; err != nil {
				return nil, err
			}
			paramValues = append(paramValues, &Ident{C: paramValue, Token: token})
			switch p.s.ch {
			case ',':
			default:
				break OUTER
			}
		}
		param.ParamValues = paramValues
		return param, p.s.err
	default:
		return nil, fmt.Errorf(`unexpected error: should got = but %s`, string(p.s.ch))
	}
}
