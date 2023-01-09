package bench

func CompareLen(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func CompareStr(s string) bool {
	if s == "" {
		return true
	}
	return false
}
