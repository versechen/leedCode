package main

import "fmt"

func isDulplicate(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return true
		} else {
			seen[c] = true
			continue

		}
	}
	return false
}
func main() {
	var s string = "ijiigehhgeg556252jifegueeuhhd123456789"
	var maxLen int
	strlen := len(s)
	var i, j int
	for i, j = 0, 1; i < strlen && j < strlen; {
		if isDulplicate(s[i : j+1]) {
			if j-i > maxLen {
				maxLen = j - i

			}
			i = j
			j++
		} else {
			j++
		}
	}

	if j-i > maxLen {
		maxLen = j - i

	}

	fmt.Printf("max len: %d", maxLen)

}
