package util

import (
	"os"
	"strings"
)

func PathDir(s string) string {
	ss := strings.Split(s, string(os.PathSeparator))
	if !strings.Contains(s, string(os.PathSeparator)) {
		return "."
	}
	return strings.Join(ss[:len(ss)-1], string(os.PathSeparator))
}
