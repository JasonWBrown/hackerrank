package main

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	level := 0
	mountains := 0
	valleys := 0
	for _, r := range path {
		previous := level
		if string(r) == "U" {
			level++
			if previous == 0 {
				mountains++
			}
		}
		if string(r) == "D" {
			level--
			if previous == 0 {
				valleys++
			}
		}

	}
	return int32(valleys)
}

func main() {}
