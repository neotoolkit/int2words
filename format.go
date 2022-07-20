package int2words

import (
	"strings"
	"unicode"
)

func WithUpperCamelCaseFormat() Option {
	return WithFormat(func(words []string) string {
		for i, v := range words {
			word := []rune(v)
			word[0] = unicode.ToTitle(word[0])
			words[i] = string(word)
		}
		return strings.Join(words, "")
	})
}

func WithLowerCamelCaseFormat() Option {
	return WithFormat(func(words []string) string {
		for i, v := range words {
			if i == 0 {
				continue
			}
			word := []rune(v)
			word[0] = unicode.ToTitle(word[0])
			words[i] = string(word)
		}
		return strings.Join(words, "")
	})
}

func WithSnakeCaseFormat() Option {
	return WithFormat(func(words []string) string {
		return strings.Join(words, "_")
	})
}

func WithKebabCaseFormat() Option {
	return WithFormat(func(words []string) string {
		return strings.Join(words, "-")
	})
}
