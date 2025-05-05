package reload

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Hex(s string) string {
	hex := regexp.MustCompile(`(\w+) \( *hex *\)`)
	res := hex.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := hex.FindStringSubmatch(s)
		v, e := strconv.ParseInt(sub_matches[1], 16, 64)
		if e != nil {
			return ""
		}
		return strconv.Itoa(int(v))
	})
	return res
}

func Bin(s string) string {
	bin := regexp.MustCompile(`(\w+) \( *bin *\)`)
	res := bin.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := bin.FindStringSubmatch(s)
		v, e := strconv.ParseInt(sub_matches[1], 2, 64)
		if e != nil {
			return ""
		}
		return strconv.Itoa(int(v))
	})
	return res
}

func Up(s string) string {
	up := regexp.MustCompile(`(\w+) \( *up *\)`)
	res := up.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := up.FindStringSubmatch(s)
		return Capitalize(sub_matches[1])
	})
	return res
}

func Low(s string) string {
	low := regexp.MustCompile(`(\w+) \( *low *\)`)
	res := low.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := low.FindStringSubmatch(s)
		return strings.ToLower(sub_matches[1])
	})
	return res
}

func Cap(s string) string {
	cap := regexp.MustCompile(`(\w+) \( *cap *\)`)
	res := cap.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := cap.FindStringSubmatch(s)
		return strings.Title(strings.ToLower(sub_matches[1]))
	})
	return res
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func All_commands(s string) string {
	res := Hex(s)
	res = Bin(res)
	res = Up(res)
	res = Low(res)
	res = Cap(res)
	return res
}
