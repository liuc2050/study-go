package comma

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	if len(s) == 0 {
		return ""
	}
	var buf bytes.Buffer
	i := len(s) % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		fmt.Fprintf(&buf, ",%s", s[i:i+3])
	}
	return buf.String()
}

func commaFloat(s string) string {
	var buf bytes.Buffer
	var begin int
	var skip bool
	for i := 0; i < len(s); i++ {
		var needWrite bool
		switch s[i] {
		case '+':
			needWrite = true
		case '-':
			needWrite = true
		case '.':
			needWrite = true
		case 'e':
			needWrite = true
			skip = false
		case 'E':
			needWrite = true
			skip = false
		}

		if skip {
			buf.WriteByte(s[i])
			begin = i + 1
			continue
		}
		if needWrite {
			if begin < i {
				buf.WriteString(comma(s[begin:i]))
			}
			buf.WriteByte(s[i])
			begin = i + 1
			if s[i] == '.' {
				skip = true
			}

		}
	}
	if begin < len(s) {
		buf.WriteString(comma(s[begin:]))
	}
	return buf.String()
}
