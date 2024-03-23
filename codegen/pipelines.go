package codegen

import (
	"strings"
	"unicode"
)

// Copied golint since it's not exported
// https://github.com/golang/lint/blob/master/lint.go#L767
//
// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func capitalize(in string) (string, error) {
	runes := []rune(in)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes), nil
}

func initialism(in string) (string, error) {
	if commonInitialisms[strings.ToUpper(in)] {
		return strings.ToUpper(in), nil
	}
	return in, nil
}

func stripPrefix(s, prefix string) (string, error) {
	return strings.TrimSpace(strings.TrimPrefix(s, prefix)), nil
}
