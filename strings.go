package utils

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func IsAllNumber(strs []string) bool {
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

func ReadLines(rd io.Reader) ([]string, error) {
	ret := []string{}
	reader := bufio.NewReader(rd)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, errors.WithStack(err)
			}
		}
		ret = append(ret, string(line))
	}
	return ret, nil
}

func DiffSubStrings(strs []string) []string {
	if len(strs) == 0 {
		return nil
	}
	prefix := ""
	suffix := ""

	idx := 0
	for {
		isEnd := false
		isAllByteSame := true
		r := byte(0)
		for _, str := range strs {
			if idx == len(str) {
				isEnd = true
				break
			}
			if r == 0 {
				r = str[idx]
			}
			if str[idx] != r {
				isEnd = true
				isAllByteSame = false
				break
			}
		}
		if isAllByteSame {
			prefix = prefix + string(r)
		}
		if isEnd {
			break
		}
		idx++
	}

	idx = 0
	for {
		isEnd := false
		isAllByteSame := true
		r := byte(0)
		for _, str := range strs {
			if len(str)-1-idx < 0 {
				isEnd = true
				break
			}
			if r == 0 {
				r = str[len(str)-1-idx]
			}
			if str[len(str)-1-idx] != r {
				isEnd = true
				isAllByteSame = false
				break
			}
		}
		if isAllByteSame {
			suffix = string(r) + suffix
		}
		if isEnd {
			break
		}
		idx++
	}
	ret := []string{}
	for _, str := range strs {
		str = strings.TrimPrefix(str, prefix)
		str = strings.TrimSuffix(str, suffix)
		ret = append(ret, str)
	}
	return ret
}
