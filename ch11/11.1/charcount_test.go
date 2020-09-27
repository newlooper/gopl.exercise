package charcount

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[rune]int
		cats    map[string]int
		utflen  map[int]int
		invalid int
	}{
		{
			input: "Say: 落红不是无情物",
			counts: map[rune]int{
				':': 1,
				'红': 1,
				'物': 1,
				'y': 1,
				' ': 1,
				'落': 1,
				'不': 1,
				'无': 1,
				'情': 1,
				'S': 1,
				'a': 1,
				'是': 1,
			},
			cats: map[string]int{
				"Unified_Ideograph":    7,
				"Ideographic":          7,
				"ASCII_Hex_Digit":      1,
				"Hex_Digit":            1,
				"Pattern_Syntax":       1,
				"Terminal_Punctuation": 1,
				"Pattern_White_Space":  1,
				"White_Space":          1,
			},
			utflen:  map[int]int{1: 5, 3: 7},
			invalid: 0,
		},
	}

	for _, test := range tests {
		counts, cats, utflen, invalid := charCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("%q counts: got %v, want %v", test.input, counts, test.counts)
		}
		if !reflect.DeepEqual(cats, test.cats) {
			t.Errorf("%q cats: got %v, want %v", test.input, cats, test.cats)
		}
		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("%q utflen: got %v, want %v", test.input, utflen, test.utflen)
		}
		if invalid != test.invalid {
			t.Errorf("%q invalid: got %v, want %v", test.input, invalid, test.invalid)
		}
	}
}
