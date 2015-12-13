package icalparser

import (
	"testing"
)

func TestFoldingWriter(t *testing.T) {
	{
		w := NewFoldingWriter(5)
		w.Write([]byte("abc"))
		w.Write([]byte("defg"))
		if e, g := "abcd\r\n efg", w.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}

	{
		w := NewFoldingWriter(5)
		w.Write([]byte("abcd"))
		if e, g := "abcd", w.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}

	{
		w := NewFoldingWriter(5)
		w.Write([]byte("abcdefghi"))
		if e, g := "abcd\r\n efgh\r\n i", w.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}

	{
		w := NewFoldingWriter(5)
		w.Write([]byte("世界"))
		if e, g := "世\r\n 界", w.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}
}
