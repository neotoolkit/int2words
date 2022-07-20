package int2words_test

import (
	"math"
	"testing"

	"github.com/neotoolkit/int2words"
)

func TestWithUpperCamelCaseFormat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		integer int64
		want    string
	}{
		{
			integer: -1,
			want:    "MinusOne",
		},
		{
			integer: 0,
			want:    "Zero",
		},
		{
			integer: 1,
			want:    "One",
		},
		{
			integer: math.MinInt64,
			want:    "MinusNineQuintillionTwoHundredTwentyThreeQuadrillionThreeHundredSeventyTwoTrillionThirtySixBillionEightHundredFiftyFourMillionSevenHundredSeventyFiveThousandEightHundredSeven",
		},
	} {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			got := int2words.Convert(tc.integer, int2words.WithUpperCamelCaseFormat())
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestWithLowerCamelCaseFormat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		integer int64
		want    string
	}{
		{
			integer: -1,
			want:    "minusOne",
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
			integer: math.MinInt64,
			want:    "minusNineQuintillionTwoHundredTwentyThreeQuadrillionThreeHundredSeventyTwoTrillionThirtySixBillionEightHundredFiftyFourMillionSevenHundredSeventyFiveThousandEightHundredSeven",
		},
	} {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			got := int2words.Convert(tc.integer, int2words.WithLowerCamelCaseFormat())
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestWithSnakeCaseFormat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		integer int64
		want    string
	}{
		{
			integer: -1,
			want:    "minus_one",
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
			integer: math.MinInt64,
			want:    "minus_nine_quintillion_two_hundred_twenty_three_quadrillion_three_hundred_seventy_two_trillion_thirty_six_billion_eight_hundred_fifty_four_million_seven_hundred_seventy_five_thousand_eight_hundred_seven",
		},
	} {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			got := int2words.Convert(tc.integer, int2words.WithSnakeCaseFormat())
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestWithKebabCaseFormat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		integer int64
		want    string
	}{
		{
			integer: -1,
			want:    "minus-one",
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
			integer: math.MinInt64,
			want:    "minus-nine-quintillion-two-hundred-twenty-three-quadrillion-three-hundred-seventy-two-trillion-thirty-six-billion-eight-hundred-fifty-four-million-seven-hundred-seventy-five-thousand-eight-hundred-seven",
		},
	} {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			got := int2words.Convert(tc.integer, int2words.WithKebabCaseFormat())
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
