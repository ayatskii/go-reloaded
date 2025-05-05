package reload

import (
	"regexp"
)

var (
	spaces       = regexp.MustCompile(` {2,}`)
	single_quote = regexp.MustCompile(`' *(([,.;:!?\w]+)( [,.;:!?\w-]+)*) *'`)
	punctuations = regexp.MustCompile(`([\w,.!?;:)']+) ([,.!?;:]+)`)
	punc_after   = regexp.MustCompile(`([,.:;?!])(\w)`)
	articles     = regexp.MustCompile(`([aA]) ['(]*([aeiouhAEIOUH]\w+)[')]*`)
)

func Clean_text(s string) string {
	res := Delete_spaces(s)
	for punctuations.Find([]byte(res)) != nil {
		res = Correct_punctuation(res)
	}
	res = Clear_single_quotes(res)
	res = Handle_articles(res)
	return res
}

func Delete_spaces(s string) string {
	res := spaces.ReplaceAllString(s, " ")
	return res
}

func Clear_single_quotes(s string) string {
	res := single_quote.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := single_quote.FindStringSubmatch(s)
		return "'" + sub_matches[1] + "'"
	})
	return res
}

func Correct_punctuation(s string) string {
	res := punctuations.ReplaceAllStringFunc(s, func(s string) string {
		spa := regexp.MustCompile(` `)
		return spa.ReplaceAllString(s, "")
	})
	res = punc_after.ReplaceAllStringFunc(res, func(s string) string {
		sub_matches := punc_after.FindStringSubmatch(s)
		return sub_matches[1] + " " + sub_matches[2]
	})
	return res
}

func Handle_articles(s string) string {
	res := articles.ReplaceAllStringFunc(s, func(str string) string {
		sub_matches := articles.FindStringSubmatch(str)
		if sub_matches[1] == "A" {
			return "An " + sub_matches[2]
		} else {
			return "an " + sub_matches[2]
		}
	})
	return res
}
