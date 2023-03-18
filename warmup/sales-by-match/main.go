package main

func main() {}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	if n < 2 {
		return 0
	}

	m := make(map[int32]int) // map[color] = count
	count := 0
	for i := 0; i < int(n); i++ {
		m[ar[i]]++
		if (m[ar[i]] > 1) && (m[ar[i]]%2 == 0) {
			count++
		}
	}

	return int32(count)
}
