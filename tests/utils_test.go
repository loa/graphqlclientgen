package tests

func stringPointerValue(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}
