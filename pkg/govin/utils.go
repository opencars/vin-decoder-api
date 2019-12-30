package govin

import (
	"regexp"
	"strings"
)

func Normalize(number string) string {
	return strings.ReplaceAll(strings.ToUpper(number), "-", "")
}

func Valid(number string) bool {
	value := Normalize(number)

	matched, err := regexp.MatchString(`^[A-HJ-NPR-Z0-9]{17}$`, value)
	if err != nil {
		panic(err)
	}

	return matched
}
