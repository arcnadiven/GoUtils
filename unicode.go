package utils

import (
	"strconv"
	"unicode"
)

func IsAllNumbers(strs []string) bool {
	for _, str := range strs {
		if !IsNumber(str) {
			return false
		}
	}
	return true
}

func IsNumber(str string) bool {
	if len(str) == 0 {
		return false
	}
	if _, err := strconv.ParseFloat(str, 64); err == nil {
		return true
	}
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return true
	}
	allNumber := true
	for _, char := range str {
		if !unicode.IsNumber(char) {
			allNumber = false
			break
		}
	}
	return allNumber
}

func HasHan(str string) bool {
	for _, char := range str {
		if unicode.Is(unicode.Han, char) {
			return true
		}
	}
	return false
}

func HasLatin(str string) bool {
	for _, char := range str {
		if unicode.Is(unicode.Latin, char) {
			return true
		}
	}
	return false
}
