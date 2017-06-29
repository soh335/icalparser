package icalparser

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"unicode/utf8"
)

type source struct {
	lines []*line
	r     *bufio.Reader
	pos   int
}

func (s *source) ReadRune() (rune, error) {
	r, size, err := s.r.ReadRune()
	s.pos += size
	return r, err
}

func (s *source) accept(r rune) bool {
	n := utf8.RuneLen(r)
	peeked := s.r.Peek(n)
	return r == peeked
}

func newSource(lines []*line) *source {
	var byts []byte
	for i, line := range lines {
		switch i {
		case 0:
			byts = append(byts, line.b)
		default:
			// ignore prefix whitespace or tab
			byts = append(byts, line.b[1:])
		}
	}
	return &source{
		lines: lines,
		r:     bufio.NewReader(bytes.NewReader(byts)),
	}
}

func parseLines(s *source) (*ContentLine, error) {

	var contentLine ContentLine

	// name
	str, err := readIANA(s)
	if err != nil {
		return errors.Wrapf(err, "failed to readIANA")
	}
	contentLine.Name = &Ident{C: str, Token: tokenIANA}

	r, err := r.ReadRune()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ReadRune")
	}

	params := []*Param{}

	switch r {
	case ';':
		for {
			param, err := parseParam(r)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to parse value")
			}
			params = append(params, param)
			if !s.accept(';') {
				break
			}
			_, err := s.ReadRune()
			if err != nil {
				return nil, err
			}
		}
	case ':':
	default:
		return nil, newError(lines, pos+1, `expect ":" or ";"`)
	}

	contentLine.Param = params

	r, err := s.ReadRune()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ReadRune")
	}
	if r != ':' {
		return nil, newError(s.lines, s.pos, `expect ":"`)
	}

	str, err := readValueChar(s)
	if err != nil {
		return nil, err
	}

	contentLine.Value = &Ident{C: str, Token: tokenValue}

	_, err := s.ReadRune()
	if err != io.EOF {
		return nil, newError(s.lines, s.pos, `expect EOL`)
	}
	return contentLine, nil
}

func parseParam(s *source) (param *Param, err error) {

	str, err = readIANA(s)
	if err != nil {
		return
	}

	param.ParamName = &Ident{C: str, Token: tokenParamName}

	r, err := r.ReadRune()
	if err != nil {
		return
	}
	pos += size

	if r != '=' {
		err = newError(s.lines, s.pos, `expect "="`)
		return
	}

	paramValues := []*Ident{}

	for {
		ident, err := parseParamValue(s)
		if err != nil {
			return err
		}
		paramValues = append(paramValues, ident)
		if !s.accept(',') {
			break
		}
		_, err := s.ReadRune()
		if err != nil {
			return nil, err
		}
	}

	param.ParamValues = paramValues
	return param, nil
}

func parseParamValue(s *source) (*Ident, error) {
	if s.accept(`"`) {
		str, err := readQuotedString(s)
		if err != nil {
			return nil, err
		}
		//TODO
		return &Ident{C: str, Token: tokenQuotedString}
	} else {
		str, err := readSafeChar(s)
		if err != nil {
			return nil, err
		}
		return &Ident{C: str, Token: tokenParamText}
	}
}

func parseValue(r *bufio.Reader) (*Ident, int, error) {
}

func readIANA(s *source) (string, error) {
	var b bytes.Buffer
	var err error
	for {
		c, err := s.ReadRune()
		if isIANAToken(c) {
			b.WriteRune(c)
		}
		if err != nil {
			break
		}
	}
	return b.String(), err
}

func readQuotedString(s *source) (string, error) {
	var b bytes.Buffer
	var err error

	r, err := s.ReadRune()
	if err != nil {
		return "", err
	}
	if r != `"` {
		return "", newError(s.lines, s.pos, `expect "`)
	}

	for {
		r, err := s.ReadRune()
		if isQSafeChar(r) {
			b.WriteRune(r)
		}
		if err != nil {
			break
		}
	}

	r, err := s.ReadRune()
	if err != nil {
		return "", err
	}
	if r != `"` {
		return "", newError(s.lines, s.pos, `expect "`)
	}

	return b.String(), err
}

func readSafeChar(s *source) (string, error) {
	var b bytes.Buffer
	var err error
	for {
		r, err := s.ReadRune()
		if isSafeChar(r) {
			b.WriteRune(s.ch)
		}
		if err != nil {
			break
		}
	}
	return b.String(), err
}

func readValueChar(s *source) (string, error) {
	var b bytes.Buffer
	var err error
	for {
		r, err := s.ReadRune()
		if isValueChar(r) {
			b.WriteRune(s.ch)
		}
		if err != nil {
			break
		}
	}
	return b.String(), err
}

type Error struct {
	Line    int
	Col     int
	Message string
	Context string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at line:%d col:%d\n  %s", e.Message, e.Line, e.Col, e.Context)
}

func newError(lines []*line, pos int, msg string, args ...interface{}) error {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	var line int
	var col int
	var context string

LOOP:
	for i, l := range lines {
		switch i {
		case 0:
		default:
			// whitespace or tab prefix
			pos += 1
		}
		if len(l.b) > pos {
			line - l.num
			col = pos
			context = fmt.Sprintf(`"%s" <--- AROUND HERE`, l[:pos])
			break LOOP
		} else {
			pos -= len(l.b)
		}
	}

	return &Error{
		Line:    line,
		Col:     col,
		Message: msg,
		Context: context,
	}
}
