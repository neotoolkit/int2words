package int2words

import (
	"math"
	"strings"
)

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type integer interface {
	signed | unsigned
}

func Convert[T integer](integer T, opts ...Option) string {
	options := Options{
		dict: Dict{
			Zero:    "zero",
			Minus:   "minus",
			Hundred: "hundred",
			Mega:    []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"},
			Unit:    []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"},
			Ten:     []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"},
			Teen:    []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
		},
		format: func(words []string) string {
			return strings.Join(words, " ")
		},
	}

	for _, opt := range opts {
		opt(&options)
	}

	if integer == 0 {
		return options.dict.Zero
	}

	var res []string

	if integer < 0 {
		res = append(res, options.dict.Minus)
		integer = T(math.Abs(float64(integer)))
	}

	var triplets []T

	for integer > 0 {
		triplets = append(triplets, T(int64(integer)%1000))
		integer = T(int64(integer) / 1000)
	}

	for i := len(triplets) - 1; i >= 0; i-- {
		triplet := triplets[i]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds > 0 {
			res = append(res, options.dict.Unit[hundreds], options.dict.Hundred)
		}

		switch tens {
		case 0:
			res = append(res, options.dict.Unit[units])
		case 1:
			res = append(res, options.dict.Teen[units])
			break
		default:
			if units > 0 {
				res = append(res, options.dict.Ten[tens], options.dict.Unit[units])
			} else {
				res = append(res, options.dict.Ten[tens])
			}
			break
		}

		if mega := options.dict.Mega[i]; mega != "" {
			res = append(res, options.dict.Mega[i])
		}
	}

	return options.format(res)
}

type Option func(opts *Options)

type Options struct {
	dict   Dict
	format func(words []string) string
}

func SetDict(d Dict) Option {
	return func(opts *Options) {
		if d.Mega[0] != "" {
			d.Mega = append(d.Mega[:1], d.Mega[0:]...)
			d.Mega[0] = ""
		}
		if d.Unit[0] != "" {
			d.Unit = append(d.Unit[:1], d.Unit[0:]...)
			d.Unit[0] = ""
		}
		if d.Ten[0] != "" {
			d.Ten = append(d.Ten[:1], d.Ten[0:]...)
			d.Ten[0] = ""
		}
		opts.dict = d
	}
}

type Dict struct {
	Zero    string
	Minus   string
	Hundred string
	Mega    []string
	Unit    []string
	Ten     []string
	Teen    []string
}

func SetFormat(format func(words []string) string) Option {
	return func(opts *Options) {
		opts.format = format
	}
}
