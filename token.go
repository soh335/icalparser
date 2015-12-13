package icalparser

type token int

const (
	tokenEOF token = iota

	tokenName
	tokenParam
	tokenValue
	tokenIANA
	tokenXName
	tokenVendorID
	tokenParamName
	tokenParamValue
	tokenParamText
	tokenQuotedString

	tokenCRLF
)
