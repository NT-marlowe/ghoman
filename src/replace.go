package main

import (
	"regexp"
	"strings"
)

// receive string and regexp, return string
func replaceAllWithRegexp(s, pattern, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, repl)
}

func trimString(line string) string {
	line = strings.ReplaceAll(line, "\t", " ")

	line = replaceAllWithRegexp(line, `\s+`, " ")
	line = replaceAllWithRegexp(line, `^#.*$`, "")
	line = replaceAllWithRegexp(line, `;\s*#`, "")

	line = strings.TrimSpace(line)

	return line
}
