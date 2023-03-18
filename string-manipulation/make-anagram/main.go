package main

import "math"

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

// a = aaaab
// b = abc
func makeAnagram(a string, b string) int32 {
	// Write your code here
	mA := createMap(a)
	mB := createMap(b)

	// mA = {a:4, b:1}
	// mB = {a:1, b:1, c:1}
	leftToDelete := 0
	for k, v := range mA {
		leftToDelete += int(math.Abs(float64(v - mB[k])))
		delete(mB, k)
	}
	for k, v := range mB {
		leftToDelete += int(math.Abs(float64(v - mA[k])))
	}

	return int32(leftToDelete)
}

func createMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func main() {}
