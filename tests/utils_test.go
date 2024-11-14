package tests

import "fmt"

func stringPointerValue(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func boolPointerValue(s *bool) string {
	if s == nil {
		return "<nil>"
	}
	return fmt.Sprint(*s)
}
