// Package bob helps you to interact with Bob
package bob

import (
	"regexp"
	"strings"
)

func isShouting(remark string) bool {
	re := regexp.MustCompile("[A-Z]")
	if re.MatchString(remark) && remark == strings.ToUpper(remark) {
		return true
	}
	return false
}

// Hey allows you to interact with Bob
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case remark[len(remark)-1] == '?' && isShouting(remark):
		return "Calm down, I know what I'm doing!"
	case isShouting(remark):
		return "Whoa, chill out!"
	case remark[len(remark)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}
