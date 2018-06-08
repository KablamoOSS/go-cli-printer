package printer

// https://github.com/willf/pad/blob/master/pad.go

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// LeftPad the string with pad up to len runes
// len may be exceeded if
func LeftPad(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

// RightPad the string with pad up to len runes
func RightPad(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}
