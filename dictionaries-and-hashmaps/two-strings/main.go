package main

func twoStrings(s1 string, s2 string) string {
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}

	for _, r := range s2 {
		if m[r] > 0 {
			return "YES"
		}
	}

	return "NO"
}

func main() {}
