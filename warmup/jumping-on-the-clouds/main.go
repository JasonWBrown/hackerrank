package main

func jumpingOnClouds(c []int32) int32 {
	jumpCount := 0
	i := 0
	end := len(c) - 1
	for i < end {
		if i+2 <= end && c[i+2] == 0 {
			jumpCount += 1
			if i+2 <= end {
				i += 2
			} else {
				i += 1
			}
			continue
		} else {
			jumpCount += 1
			i += 1
			continue
		}
	}

	return int32(jumpCount)
}

func main() {}
