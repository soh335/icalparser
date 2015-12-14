package icalparser

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrinter(t *testing.T) {
	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//hacksw/handcal//NONSGML v1.0//EN
BEGIN:VEVENT
UID:19970610T172345Z-AF23B2@example.com
DTSTAMP:19970610T172345Z
DTSTART:19970714T170000Z
DTEND:19970715T040000Z
SUMMARY:Bastille Day Party
END:VEVENT
END:VCALENDAR
`,
		)

		o, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}
		var b bytes.Buffer
		NewPrinter(o).WriteTo(&b)
		if e, g := data, b.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
ATTENDEE;DELEGATED-TO="mailto:jdoe@example.com","mailto:jqpublic@example.c
 om":mailto:jsmith@example.com
END:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}
		var b bytes.Buffer
		NewPrinter(obj).WriteTo(&b)
		if e, g := data, b.String(); e != g {
			t.Errorf("expected %v but got %v", e, g)
		}
	}
}
