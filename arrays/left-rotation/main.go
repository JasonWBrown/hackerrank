package main

func main() {}

func rotLeft(a []int32, d int32) []int32 {
	// Write your code here
	b := make([]int32, len(a))
	offset := int(d)
	for i := 0; i < len(a)-offset; i++ {
		b[i] = a[i+offset]
	}

	offset = len(a) - offset
	for i := 0; i < len(a)-offset; i++ {
		b[i+offset] = a[i]
	}

	return b
}
