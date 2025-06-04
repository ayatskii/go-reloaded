package reload

import (
	"errors"
	"regexp"
	cleaner "reload/cleaner"
	"strconv"
	"strings"
	"unicode"
)

var (
	command = regexp.MustCompile(`(?i)\( ?(hex|bin|up|low|cap)(,? (-*\d+)?)?\) *`)
)

func cap(str string) string {
	if len(str) == 0 {
		return str
	}
	first := string(str[0])
	rest := ""
	if len(str) > 1 {
		rest = strings.ToLower(str[1:])
	}
	switch first {
	case "'", "\"", "(":
		if len(str) > 1 {
			return first + strings.ToUpper(string(str[1])) + rest[1:] + " "
		}
		return first
	default:
		return strings.ToUpper(first) + rest + " "
	}
}

func replace_at_index(input string, replacement string, index int, length int) string {
	return input[:index] + replacement + input[index+length:]
}
func find_words(amount int, s string) ([]int, error) {
	var res []int
	ins := command.FindStringIndex(s)
	start := ins[0]
	for ; amount > 0; amount-- {
		start--
		if start < 0 {
			start = 0
			break
		}
		for !unicode.IsLetter(rune(s[start])) {
			start--
			if start < 0 {
				start = 0
				break
			}
		}
		res = append(res, start)
		for unicode.IsLetter(rune(s[start])) {
			start--
			if start < 0 {
				break
			}
		}
		res = append(res, start+1)
	}
	if len(res) == 0 {
		return res, errors.New("zero words")
	}
	return res, nil
}

func Proceed_commands(s string) string {
	res := s
	for command.FindStringIndex(res) != nil {
		match := command.FindStringSubmatch(res)
		amount := 1
		if match[3] != "" {
			amount, _ = strconv.Atoi(match[3])
			if amount <= 0 {
				ins := command.FindStringIndex(res)
				res = replace_at_index(res, "", ins[0], ins[1]-ins[0])
				continue
			}
		}
		words, err := find_words(amount, res)
		if err != nil {
			ins := command.FindStringIndex(res)
			res = replace_at_index(res, "", ins[0], ins[1]-ins[0])
			continue
		}
		switch strings.ToLower(match[1]) {
		case "hex":
			hex := regexp.MustCompile(`(\w+)\s*\(\s*hex\s*\)`)
			matches := hex.FindStringSubmatchIndex(res)
			if len(matches) != 4 {
				break
			}
			word, _ := strconv.ParseInt(res[matches[2]:matches[3]], 16, 64)

			res = replace_at_index(res, strconv.FormatInt(word, 10), matches[2], matches[3]-matches[2])
		case "bin":
			bin := regexp.MustCompile(`(\w+)\s*\(\s*bin\s*\)`)
			matches := bin.FindStringSubmatchIndex(res)
			if len(matches) != 4 {
				break
			}
			word, _ := strconv.ParseInt(res[matches[2]:matches[3]], 2, 64)
			res = replace_at_index(res, strconv.FormatInt(word, 10), matches[2], matches[3]-matches[2])
		case "low":
			for i := 0; i < len(words); i += 2 {
				word := strings.ToLower(res[words[i+1] : words[i]+1])

				res = replace_at_index(res, word, words[i+1], words[i]-words[i+1]+1)
			}
		case "up":
			for i := 0; i < len(words); i += 2 {
				word := strings.ToUpper(res[words[i+1] : words[i]+1])

				res = replace_at_index(res, word, words[i+1], words[i]-words[i+1]+1)
			}
		case "cap":
			for i := 0; i < len(words); i += 2 {
				word := strings.ToLower(res[words[i+1] : words[i]+1])
				word = cap(word)
				res = replace_at_index(res, word, words[i+1], words[i]-words[i+1]+1)
			}
		}
		ins := command.FindStringIndex(res)
		res = replace_at_index(res, "", ins[0], ins[1]-ins[0])
		res = cleaner.Clean_text(res)
	}
	return res
}

// func Hex(s string) string {
// 	hex := regexp.MustCompile(`(\w+) *\( *hex *\)`)
// 	res := hex.ReplaceAllStringFunc(s, func(s string) string {
// 		sub_matches := hex.FindStringSubmatch(s)
// 		v, e := strconv.ParseInt(sub_matches[1], 16, 64)
// 		if e != nil {
// 			return ""
// 		}
// 		return strconv.Itoa(int(v))
// 	})
// 	return res
// }

// func Bin(s string) string {
// 	bin := regexp.MustCompile(`(\w+) *\( *bin *\)`)
// 	res := bin.ReplaceAllStringFunc(s, func(s string) string {
// 		sub_matches := bin.FindStringSubmatch(s)
// 		v, e := strconv.ParseInt(sub_matches[1], 2, 64)
// 		if e != nil {
// 			return ""
// 		}
// 		return strconv.Itoa(int(v))
// 	})
// 	return res
// }

// func Up(s string) string {
// 	up := regexp.MustCompile(`(\w+) *\( *up *\)`)
// 	res := up.ReplaceAllStringFunc(s, func(s string) string {
// 		sub_matches := up.FindStringSubmatch(s)
// 		return strings.ToUpper(sub_matches[1])
// 	})
// 	return res
// }

// func Low(s string) string {
// 	low := regexp.MustCompile(`(\w+) *\( *low *\)`)
// 	res := low.ReplaceAllStringFunc(s, func(s string) string {
// 		sub_matches := low.FindStringSubmatch(s)
// 		return strings.ToLower(sub_matches[1])
// 	})
// 	return res
// }

// func Capitalize(s string) string {
// 	if s == "" {
// 		return ""
// 	}
// 	runes := []rune(s)
// 	runes[0] = unicode.ToUpper(runes[0])
// 	return string(runes)
// }

// func Cap(s string) string {
// 	cap := regexp.MustCompile(`(\w+) *\( *cap *\)`)
// 	res := cap.ReplaceAllStringFunc(s, func(s string) string {
// 		sub_matches := cap.FindStringSubmatch(s)
// 		return Capitalize(strings.ToLower(sub_matches[1]))
// 	})
// 	return res
// }

// func All_commands(s string) string {
// 	res := Hex(s)
// 	res = Bin(res)
// 	res = Up(res)
// 	res = Low(res)
// 	res = Cap(res)
// 	return res
// }
