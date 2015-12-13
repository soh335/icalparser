package icalparser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
SUMMARY:Bastille Day Party
END:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, &Object{
			headerLine: &ContentLine{
				name:  &Ident{"BEGIN", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
			propertiyLines: []*ContentLine{
				{
					name:  &Ident{"VERSION", tokenIANA},
					param: []*Param{},
					value: &Ident{"2.0", tokenValue},
				},
			},
			components: []*Component{
				{
					headerLine: &ContentLine{
						name:  &Ident{"BEGIN", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
					propertiyLines: []*ContentLine{
						{
							name:  &Ident{"SUMMARY", tokenIANA},
							param: []*Param{},
							value: &Ident{"Bastille Day Party", tokenValue},
						},
					},
					footerLine: &ContentLine{
						name:  &Ident{"END", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			footerLine: &ContentLine{
				name:  &Ident{"END", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
		}) {
			t.Error("got unexpected obj")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
DESCRIPTION:This is a lo
 ng description
  that exists on a long line.
END:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, &Object{
			headerLine: &ContentLine{
				name:  &Ident{"BEGIN", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
			propertiyLines: []*ContentLine{
				{
					name:  &Ident{"VERSION", tokenIANA},
					param: []*Param{},
					value: &Ident{"2.0", tokenValue},
				},
			},
			components: []*Component{
				{
					headerLine: &ContentLine{
						name:  &Ident{"BEGIN", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
					propertiyLines: []*ContentLine{
						{
							name:  &Ident{"DESCRIPTION", tokenIANA},
							param: []*Param{},
							value: &Ident{"This is a long description that exists on a long line.", tokenValue},
						},
					},
					footerLine: &ContentLine{
						name:  &Ident{"END", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			footerLine: &ContentLine{
				name:  &Ident{"END", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
		}) {
			t.Error("got unexpected obj")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
ATTENDEE;RSVP="TRUE";ROLE=REQ-PARTICIPANT:mailto:
 jsmith@example.com
END:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, &Object{
			headerLine: &ContentLine{
				name:  &Ident{"BEGIN", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
			propertiyLines: []*ContentLine{
				{
					name:  &Ident{"VERSION", tokenIANA},
					param: []*Param{},
					value: &Ident{"2.0", tokenValue},
				},
			},
			components: []*Component{
				{
					headerLine: &ContentLine{
						name:  &Ident{"BEGIN", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
					propertiyLines: []*ContentLine{
						{
							name: &Ident{"ATTENDEE", tokenIANA},
							param: []*Param{
								{
									paramName:   &Ident{"RSVP", tokenParamName},
									paramValues: []*Ident{{"TRUE", tokenQuotedString}},
								},
								{
									paramName:   &Ident{"ROLE", tokenParamName},
									paramValues: []*Ident{{"REQ-PARTICIPANT", tokenParamText}},
								},
							},
							value: &Ident{"mailto:jsmith@example.com", tokenValue},
						},
					},
					footerLine: &ContentLine{
						name:  &Ident{"END", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			footerLine: &ContentLine{
				name:  &Ident{"END", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
		}) {
			t.Error("got unexpected obj")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
ATTENDEE;DELEGATED-TO="mailto:jdoe@example.com","mailto:jqpublic
 @example.com":mailto:jsmith@example.com
END:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, &Object{
			headerLine: &ContentLine{
				name:  &Ident{"BEGIN", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
			propertiyLines: []*ContentLine{
				{
					name:  &Ident{"VERSION", tokenIANA},
					param: []*Param{},
					value: &Ident{"2.0", tokenValue},
				},
			},
			components: []*Component{
				{
					headerLine: &ContentLine{
						name:  &Ident{"BEGIN", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
					propertiyLines: []*ContentLine{
						{
							name: &Ident{"ATTENDEE", tokenIANA},
							param: []*Param{
								{
									paramName: &Ident{"DELEGATED-TO", tokenParamName},
									paramValues: []*Ident{
										{"mailto:jdoe@example.com", tokenQuotedString},
										{"mailto:jqpublic@example.com", tokenQuotedString},
									},
								},
							},
							value: &Ident{"mailto:jsmith@example.com", tokenValue},
						},
					},
					footerLine: &ContentLine{
						name:  &Ident{"END", tokenIANA},
						param: []*Param{},
						value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			footerLine: &ContentLine{
				name:  &Ident{"END", tokenIANA},
				param: []*Param{},
				value: &Ident{"VCALENDAR", tokenValue},
			},
		}) {
			t.Error("got unexpected obj")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
SUMMARY:Bastille Day Party
END:VEVENT
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()

		if err == nil {
			t.Error("should got err")
		}

		if obj != nil {
			t.Error("obj should be nil")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
SUMMARY:Bastille Day Party
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()

		if err == nil {
			t.Error("should got err")
		}

		if obj != nil {
			t.Error("obj should be nil")
		}
	}

	{
		data := rnString(`BEGIN:VCALENDAR
VERSION:2.0
BEGIN:VEVENT
SUMMARY:Bastille Day Party
E,ND:VEVENT
END:VCALENDAR
`)

		obj, err := NewParser(strings.NewReader(data)).Parse()

		if err == nil {
			t.Error("should got err")
		}

		if obj != nil {
			t.Error("obj should be nil")
		}
	}
}

func rnString(s string) string {
	return strings.Replace(s, "\n", "\r\n", -1)
}
