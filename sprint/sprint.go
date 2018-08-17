package sprint

import "strconv"

type FileInt int

func Sprint(v interface{}) string {
	switch v := v.(type) {
	case int:
		return "int " + strconv.Itoa(v)
	case FileInt:
		return "FileInt " + strconv.Itoa(int(v))
	default:
		return "???"
	}
}
