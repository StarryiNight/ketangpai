package logic

import (
	"unicode"
	"unicode/utf8"
)

// TruncateByWords 把文章内容截取到限制的最大长度
func TruncateByWords(s string, maxWords int) string {
	processedWords := 0
	wordStarted := false
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRuneInString(s[i:])
		if !isSeparator(r) {
			i += width
			wordStarted = true
			continue
		}

		if !wordStarted {
			i += width
			continue
		}

		wordStarted = false
		processedWords++
		if processedWords == maxWords {
			const ending = "..."
			if (i + len(ending)) >= len(s) {
				// Source string ending is shorter than "..."
				return s
			}

			return s[:i] + ending
		}

		i += width
	}


	return s
}

func isSeparator(r rune) bool {
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}

	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	return unicode.IsSpace(r)
}
