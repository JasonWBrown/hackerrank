package main

func main() {}

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

// not completely happy with this algo but it gets it done.
func sherlockAndAnagrams(s string) int32 {
	numOfAnangramSubstrings := int32(0)

	//i is the start index
	for i := 0; i <= len(s)-1; i++ {
		//j is the start of the substring
		for j := i + 1; j <= len(s)-1; j++ {
			// k is the length of the string
			for k := 1; k <= len(s)-j; k++ {
				if isAnagram(s[i:i+k], s[j:j+k]) {
					numOfAnangramSubstrings++
				}
				// fmt.Printf("start index %d substring index %d length %d the substrings <<%s>> <<%s>> number %d len %d\n", i, j, k, s[i:i+k], s[j:j+k], numOfAnangramSubstrings, len(s))
			}
		}
	}

	return numOfAnangramSubstrings
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	mS1 := make(map[rune]int)
	for _, r := range s1 {
		mS1[r]++
	}
	for _, r := range s2 {
		if mS1[r] == 0 {
			return false
		} else {
			mS1[r]--
		}
	}
	for _, v := range mS1 {
		if v > 0 {
			return false
		}
	}
	return true
}
