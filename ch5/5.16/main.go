package main

import "bytes"

func main() {
	println(join("|*|", "Life", "若", "Only", "如", "First Meet"))
}

func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var b bytes.Buffer
	for _, s := range strs {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.Truncate(b.Len() - len(sep))
	return b.String()
}
