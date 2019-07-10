package util

import (
	"bytes"
)

func StringConcat(str ...string) string {
	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}
