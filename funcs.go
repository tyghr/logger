package logger

func TruncateString(s string, l int) string {
	if len(s) > l {
		return s[0:l]
	}
	return s
}
