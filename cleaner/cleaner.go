package reload

import (
	"regexp"
)

var (
	spaces          = regexp.MustCompile(` {2,}`)
	single_quote    = regexp.MustCompile(`' *(([,.;:!?\w]+)( [,.;:!?\w-]+)*) *'`)
	punctuations    = regexp.MustCompile(`([\w,.!?;:)']+) ([,.!?;:]+)`)
	punc_after      = regexp.MustCompile(`([,.:;?!])(\w)`)
	articles        = regexp.MustCompile(`(([ |\n])?(?:\s|^))([aA]) ((['"]|\( ?)?([aeiouhAEIOUH]\w+)( ?\)|['")])?)`)
	double_quote    = regexp.MustCompile(`" *(([,.;:!?\w]+)( [,.;:!?\w-]+)*) *"`)
	scopes          = regexp.MustCompile(`\( *(([,.;:!?\w]+)( [,.;:!?\w-]+)*) *\)`)
	delimeter       = regexp.MustCompile(`['"(]\w+['")]`)
	delimetrs_empty = regexp.MustCompile(`''|""`)
	new_line        = regexp.MustCompile(`(\n) +`)
)

func Clean_text(s string) string {
	res := delete_spaces(s)
	res = correct_punctuation(res)
	res = clear_single_quotes(res)
	res = clear_double_quotes(res)
	res = clear_scopes(res)
	res = correct_punctuation(res)
	res = handle_delimeters(res)
	res = handle_empty_delimetrs(res)
	res = delete_spaces(res)
	res = handle_articles(res)
	res = handle_new_line(res)

	return res
}

func delete_spaces(s string) string {
	res := spaces.ReplaceAllString(s, " ")
	return res
}

func clear_single_quotes(s string) string {
	res := single_quote.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := single_quote.FindStringSubmatch(s)
		return "'" + sub_matches[1] + "'"
	})
	return res
}

func clear_double_quotes(s string) string {
	res := double_quote.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := double_quote.FindStringSubmatch(s)
		return "\"" + sub_matches[1] + "\""
	})
	return res
}

func clear_scopes(s string) string {
	res := scopes.ReplaceAllStringFunc(s, func(s string) string {
		sub_matches := scopes.FindStringSubmatch(s)
		return "(" + sub_matches[1] + ")"
	})
	return res
}

func correct_punctuation(s string) string {
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

func handle_delimeters(s string) string {
	res := delimeter.ReplaceAllStringFunc(s, func(str string) string {
		matche := delimeter.FindString(str)
		return " " + matche + " "
	})
	return res
}

func handle_empty_delimetrs(s string) string {
	res := delimetrs_empty.ReplaceAllStringFunc(s, func(str string) string {
		return " " + str + " "
	})
	return res
}

func handle_new_line(s string) string {
	res := new_line.ReplaceAllStringFunc(s, func(str string) string {
		return "\n"
	})
	return res
}

func handle_articles(s string) string {
	res := articles.ReplaceAllStringFunc(s, func(str string) string {
		sub_matches := articles.FindStringSubmatch(str)
		if sub_matches[6] == "hex" || sub_matches[6] == "up" {
			return sub_matches[0]
		}
		if sub_matches[3] == "A" {
			return sub_matches[1] + "An " + sub_matches[4]
		} else {
			return sub_matches[1] + "an " + sub_matches[4]
		}
	})
	return res
}
