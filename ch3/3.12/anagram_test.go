package main

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"a", "ab", false},
		{"aba", "baa", true},
		{"aaa", "baa", false},
		{"aaabbb", "abbbbb", false},
	}
	for _, test := range tests {
		got := anagram(test.a, test.b)
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
}
