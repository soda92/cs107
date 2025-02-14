package main

import (
	"strings"
)

func get_output2(file string) string {
	go_str := get_output(file)
	lines := []string{}
	s := ""

	for _, v := range go_str {
		if v == '\n' {
			continue
		}
		if len(s) > 80 && v == ' ' {
			lines = append(lines, s)
			s = ""
		} else {
			if len(s) > 1 && (v == ',' || v == '.') && s[len(s)-1] == ' ' {
				s = s[:len(s)-1]
			}
			s = s + string(v)
		}
	}
	lines = append(lines, s)
	return strings.Join(lines, string('\n'))
}
