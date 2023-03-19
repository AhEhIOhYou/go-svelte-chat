package utils

import "regexp"

func CheckAlphaNumeric(str string) bool {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	return IsAlphaNumeric(str)
}
