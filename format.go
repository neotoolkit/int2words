package int2words

import (
	"strings"
	"unicode"
)

func SetUpperCamelCaseFormat() Option {
	return func(opts *Options) {
		opts.format = func(words []string) string {
			for i, v := range words {
				word := []rune(v)
				word[0] = unicode.ToTitle(word[0])
				words[i] = string(word)
			}
			return strings.Join(words, "")
		}
	}
}

func SetLowerCamelCaseFormat() Option {
	return func(opts *Options) {
		opts.format = func(words []string) string {
			for i, v := range words {
				if i == 0 {
					continue
				}
				word := []rune(v)
				word[0] = unicode.ToTitle(word[0])
				words[i] = string(word)
			}
			return strings.Join(words, "")
		}
	}
}

func SetSnakeCaseFormat() Option {
	return func(opts *Options) {
		opts.format = func(words []string) string {
			return strings.Join(words, "_")
		}
	}
}

func SetKebabCaseFormat() Option {
	return func(opts *Options) {
		opts.format = func(words []string) string {
			return strings.Join(words, "-")
		}
	}
}
