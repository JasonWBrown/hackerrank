package main

func repeatedString(s string, n int64) int64 {
	if n == 0 {
		return 0
	}
	// get the number of runes/characters in the string
	size := len(s)
	// divided n by the number of characters to see how many times len(s) goes into n
	if size <= int(n) {
		quotient := n / int64(size)
		mod := n % int64(size)
		count := 0
		// determine the number of a's in the string
		for _, r := range s {
			if string(r) == "a" {
				count++
			}
		}
		// multiply number of a's by step 2
		whole := int64(count * int(quotient))

		count = 0
		for i, r := range s {
			if i == int(mod) {
				break
			}
			if string(r) == "a" {
				count++
			}
		}
		return whole + int64(count)
	} else {
		// range through just to n return that number
		count := 0
		for i, r := range s {
			if i == int(n) {
				return int64(count)
			}
			if string(r) == "a" {
				count++
			}
		}
	}

	return 0
}
func repeatedString2(s string, n int64) int64 {
	if n == 0 {
		return 0
	}
	// get the number of runes/characters in the string
	size := len(s)
	// divided n by the number of characters to see how many times len(s) goes into n
	if size <= int(n) {
		quotient := n / int64(size)
		count := 0
		// determine the number of a's in the string
		for _, r := range s {
			if string(r) == "a" {
				count++
			}
		}
		// multiply number of a's by step 2
		return int64(count * int(quotient))
	} else {
		// range through just to n return that number
		count := 0
		for i, r := range s {
			if i == int(n)-1 {
				return int64(count)
			}
			if string(r) == "a" {
				count++
			}
		}
	}

	return 0
}

func main() {}
