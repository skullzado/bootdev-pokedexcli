package main

import (
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	var words []string

	if text == "" {
		return words
	}

	for t := range strings.SplitSeq(strings.TrimSpace(text), " ") {
		if t == "" {
			continue
		}
		re := regexp.MustCompile(`\P{L}+`)
		var formatted string = strings.ToLower(t)
		formatted = strings.TrimSpace(formatted)
		formatted = re.ReplaceAllString(formatted, "")
		words = append(words, formatted)
	}

	return words
}
