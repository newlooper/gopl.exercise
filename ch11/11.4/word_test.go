package word

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

var punctuation = []rune{' '}

func init() {
	for r := rune(0x21); r < 0x7e; r++ {
		if unicode.IsPunct(r) {
			punctuation = append(punctuation, r)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
		p = randomPalindromeWithPunctuation(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents are derived
// from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomPalindromeWithPunctuation(rng *rand.Rand) string {
	n := 10 + rng.Intn(25) // random length up to 34
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	b := &bytes.Buffer{}
	for _, r := range runes {
		if rng.Float64() < 0.1 {
			b.WriteRune(randomPunctuation(rng))
		}
		b.WriteRune(r)
	}
	return b.String()
}

func randomPunctuation(rng *rand.Rand) rune {
	return punctuation[rng.Intn(len(punctuation))]
}
