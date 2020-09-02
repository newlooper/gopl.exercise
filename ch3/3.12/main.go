package main

func main() {

}

func anagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		if i, ok := m[r]; ok {
			if i > 1 {
				m[r]--
			} else {
				delete(m, r)
			}
		} else {
			return false
		}
	}

	return len(m) == 0
}
