package icalparser

import (
	"bufio"
	"bytes"
	"fmt"
	"unicode/utf8"
)

type scanner struct {
	r   *bufio.Reader
	ch  rune
	err error
}

func (s *scanner) read() {
	r, _, err := s.r.ReadRune()
	if err != nil {
		s.err = err
	}
	s.ch = r

	if s.err != nil {
		return
	}

	if s.accept([]rune{'\r', '\n', ' '}) || s.accept([]rune{'\r', '\n', '\t'}) {
		s.r.ReadRune()
		s.r.ReadRune()
		s.r.ReadRune()
	}

	return
}

func (s *scanner) accept(expected []rune) bool {
	rawExpected := make([]byte, 4*len(expected))
	offset := 0

	for _, r := range expected {
		n := utf8.EncodeRune(rawExpected[offset:], r)
		offset += n
	}

	rawExpected = rawExpected[:offset]

	peeked, _ := s.r.Peek(len(rawExpected))
	return bytes.Equal(peeked, rawExpected)
}

func (s *scanner) unread() error {
	return s.r.UnreadRune()
}

func (s *scanner) isCRLF() bool {
	if s.ch != '\r' {
		return false
	}

	return s.accept([]rune{'\n'})
}

func (s *scanner) scanName() (string, token) {
	var b bytes.Buffer

	for {
		s.read()
		if s.err != nil {
			return b.String(), tokenIANA
		}

		if isIANAToken(s.ch) {
			b.WriteRune(s.ch)
			continue
		}

		return b.String(), tokenIANA
	}
}

func (s *scanner) scanParamName() (string, token) {
	str, _ := s.scanName()
	return str, tokenParamName
}

func (s *scanner) scanParamValue() (string, token) {
	if s.accept([]rune{'"'}) {
		s.read()
		paramValue := s.scanQuotedString()
		switch s.ch {
		case '"':
			s.read() // " should be readed
		default:
			s.err = fmt.Errorf(`unexpected error: should got "`)
			return "", tokenQuotedString
		}

		return paramValue, tokenQuotedString
	} else {
		paramValue := s.scanParamText()
		return paramValue, tokenParamText
	}
}

func (s *scanner) scanQuotedString() string {
	var b bytes.Buffer
	for {
		s.read()
		if s.err != nil {
			return b.String()
		}
		if isQSafeChar(s.ch) {
			b.WriteRune(s.ch)
			continue
		}

		return b.String()
	}
}

func (s *scanner) scanParamText() string {
	var b bytes.Buffer
	for {
		s.read()
		if s.err != nil {
			return b.String()
		}
		if isSafeChar(s.ch) {
			b.WriteRune(s.ch)
			continue
		}

		return b.String()
	}
}

func (s *scanner) scanValue() (string, token) {
	var b bytes.Buffer
	for {
		s.read()
		if s.err != nil {
			return b.String(), tokenValue
		}
		if isValueChar(s.ch) {
			b.WriteRune(s.ch)
			continue
		}

		return b.String(), tokenValue
	}

}

func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func isDigit(r rune) bool {
	return ('0' <= r && r <= '9')
}

func isIANAToken(r rune) bool {
	return isAlpha(r) || isDigit(r) || r == '-'
}

func isSafeChar(r rune) bool {
	return r == ' ' || r == 0x21 || (0x23 <= r && r <= 0x2b) || (0x2d <= r && r <= 0x39) || (0x3c <= r && r <= 0x7e) || isNonUSAscii(r)
}

func isQSafeChar(r rune) bool {
	return r == ' ' || r == 0x21 || (0x23 <= r && r <= 0x7e) || isNonUSAscii(r)
}

func isValueChar(r rune) bool {
	return r == ' ' || (0x21 <= r && r <= 0x7e) || isNonUSAscii(r)
}

func isNonUSAscii(r rune) bool {
	// NON-US-ASCII  = UTF8-2 / UTF8-3 / UTF8-4
	// https://tools.ietf.org/html/rfc3629#section-4
	return !(0x00 <= r && r <= 0x7f) && utf8.ValidRune(r)
}
