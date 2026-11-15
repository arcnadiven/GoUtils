package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ContainsAllItems 判断 text 是否包含 items 中的所有子串
func ContainsAllItems(text string, items []string) bool {
	text = strings.TrimSpace(text)
	has := true
	for _, item := range items {
		if !strings.Contains(strings.ToLower(text), strings.ToLower(item)) {
			has = false
			break
		}
	}
	return has
}

// ContainsItem 判断 text 是否包含 items 中的一个子串
func ContainsItem(text string, items []string) (string, bool) {
	text = strings.TrimSpace(text)
	for _, item := range items {
		lowerText := strings.ToLower(text)
		lowerItem := strings.ToLower(item)
		if strings.Contains(lowerText, lowerItem) {
			return item, true
		}
	}
	return "", false
}

// HasPrefixItem 判断 text 的前缀是否匹配 items 中的一个子串
func HasPrefixItem(text string, items []string) (string, bool) {
	text = strings.TrimSpace(text)
	for _, item := range items {
		lowerText := strings.ToLower(text)
		lowerItem := strings.ToLower(item)
		if strings.HasPrefix(lowerText, lowerItem) {
			return item, true
		}
	}
	return "", false
}

// HasItem 判断 text 是否等于 items 中的一个子串
func HasItem(text string, items []string) (string, bool) {
	text = strings.TrimSpace(text)
	for _, item := range items {
		lowerText := strings.ToLower(text)
		lowerItem := strings.ToLower(item)
		if lowerText == lowerItem {
			return item, true
		}
	}
	return "", false
}

// ToTitle 转换为标题格式（首字母大写）
func ToTitle(str string) string {
	title := cases.Title(language.Und)
	return title.String(str)
}

func Int32Join(items []int32, sep string) string {
	ret := []string{}
	for _, item := range items {
		ret = append(ret, strconv.Itoa(int(item)))
	}
	return strings.Join(ret, sep)
}

func Int64Join(items []int64, sep string) string {
	ret := []string{}
	for _, item := range items {
		ret = append(ret, strconv.FormatInt(item, 10))
	}
	return strings.Join(ret, sep)
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
