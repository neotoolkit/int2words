package int2words_test

import (
	"math"
	"testing"

	"github.com/neotoolkit/int2words"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		integer int64
		want    string
		opts    []int2words.Option
	}{
		{
			integer: -1,
			want:    "minus one",
		},
		{
			integer: 0,
			want:    "zero",
		},
		{
			integer: 1,
			want:    "one",
		},
		{
			integer: 10,
			want:    "ten",
		},
		{
			integer: 110,
			want:    "one hundred ten",
		},
		{
			integer: 1110,
			want:    "one thousand one hundred ten",
		},
		{
			integer: 11110,
			want:    "eleven thousand one hundred ten",
		},
		{
			integer: 111110,
			want:    "one hundred eleven thousand one hundred ten",
		},
		{
			integer: 1111110,
			want:    "one million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 11111110,
			want:    "eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 111111110,
			want:    "one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 1111111110,
			want:    "one billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 11111111110,
			want:    "eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 111111111110,
			want:    "one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 1111111111110,
			want:    "one trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 11111111111110,
			want:    "eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 111111111111110,
			want:    "one hundred eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 1111111111111110,
			want:    "one quadrillion one hundred eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 11111111111111110,
			want:    "eleven quadrillion one hundred eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 111111111111111110,
			want:    "one hundred eleven quadrillion one hundred eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: 1111111111111111110,
			want:    "one quintillion one hundred eleven quadrillion one hundred eleven trillion one hundred eleven billion one hundred eleven million one hundred eleven thousand one hundred ten",
		},
		{
			integer: math.MinInt64,
			want:    "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven",
		},
		{
			integer: math.MaxInt64,
			want:    "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven",
		},
	} {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			got := int2words.Convert(tc.integer, tc.opts...)
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
