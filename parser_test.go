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
			HeaderLine: &ContentLine{
				Name:  &Ident{"BEGIN", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
			},
			PropertiyLines: []*ContentLine{
				{
					Name:  &Ident{"VERSION", tokenIANA},
					Param: []*Param{},
					Value: &Ident{"2.0", tokenValue},
				},
			},
			Components: []*Component{
				{
					HeaderLine: &ContentLine{
						Name:  &Ident{"BEGIN", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
					PropertiyLines: []*ContentLine{
						{
							Name:  &Ident{"SUMMARY", tokenIANA},
							Param: []*Param{},
							Value: &Ident{"Bastille Day Party", tokenValue},
						},
					},
					FooterLine: &ContentLine{
						Name:  &Ident{"END", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			FooterLine: &ContentLine{
				Name:  &Ident{"END", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
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
			HeaderLine: &ContentLine{
				Name:  &Ident{"BEGIN", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
			},
			PropertiyLines: []*ContentLine{
				{
					Name:  &Ident{"VERSION", tokenIANA},
					Param: []*Param{},
					Value: &Ident{"2.0", tokenValue},
				},
			},
			Components: []*Component{
				{
					HeaderLine: &ContentLine{
						Name:  &Ident{"BEGIN", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
					PropertiyLines: []*ContentLine{
						{
							Name:  &Ident{"DESCRIPTION", tokenIANA},
							Param: []*Param{},
							Value: &Ident{"This is a long description that exists on a long line.", tokenValue},
						},
					},
					FooterLine: &ContentLine{
						Name:  &Ident{"END", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			FooterLine: &ContentLine{
				Name:  &Ident{"END", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
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
			HeaderLine: &ContentLine{
				Name:  &Ident{"BEGIN", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
			},
			PropertiyLines: []*ContentLine{
				{
					Name:  &Ident{"VERSION", tokenIANA},
					Param: []*Param{},
					Value: &Ident{"2.0", tokenValue},
				},
			},
			Components: []*Component{
				{
					HeaderLine: &ContentLine{
						Name:  &Ident{"BEGIN", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
					PropertiyLines: []*ContentLine{
						{
							Name: &Ident{"ATTENDEE", tokenIANA},
							Param: []*Param{
								{
									ParamName:   &Ident{"RSVP", tokenParamName},
									ParamValues: []*Ident{{"TRUE", tokenQuotedString}},
								},
								{
									ParamName:   &Ident{"ROLE", tokenParamName},
									ParamValues: []*Ident{{"REQ-PARTICIPANT", tokenParamText}},
								},
							},
							Value: &Ident{"mailto:jsmith@example.com", tokenValue},
						},
					},
					FooterLine: &ContentLine{
						Name:  &Ident{"END", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			FooterLine: &ContentLine{
				Name:  &Ident{"END", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
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
			HeaderLine: &ContentLine{
				Name:  &Ident{"BEGIN", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
			},
			PropertiyLines: []*ContentLine{
				{
					Name:  &Ident{"VERSION", tokenIANA},
					Param: []*Param{},
					Value: &Ident{"2.0", tokenValue},
				},
			},
			Components: []*Component{
				{
					HeaderLine: &ContentLine{
						Name:  &Ident{"BEGIN", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
					PropertiyLines: []*ContentLine{
						{
							Name: &Ident{"ATTENDEE", tokenIANA},
							Param: []*Param{
								{
									ParamName: &Ident{"DELEGATED-TO", tokenParamName},
									ParamValues: []*Ident{
										{"mailto:jdoe@example.com", tokenQuotedString},
										{"mailto:jqpublic@example.com", tokenQuotedString},
									},
								},
							},
							Value: &Ident{"mailto:jsmith@example.com", tokenValue},
						},
					},
					FooterLine: &ContentLine{
						Name:  &Ident{"END", tokenIANA},
						Param: []*Param{},
						Value: &Ident{"VEVENT", tokenValue},
					},
				},
			},
			FooterLine: &ContentLine{
				Name:  &Ident{"END", tokenIANA},
				Param: []*Param{},
				Value: &Ident{"VCALENDAR", tokenValue},
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
