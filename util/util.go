package util

func ToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
