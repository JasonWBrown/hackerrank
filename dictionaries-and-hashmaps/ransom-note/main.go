package main

func main() {}

func checkMagazine(magazine []string, note []string) bool {
	// we know that if the number of words in the note is greater than number of words in magazine
	// we can not create the ransom note
	if len(note) > len(magazine) {
		return false
	}

	// iterate through all of the magazine and create a map
	// map key is the word value is the number of times it occurs
	magMap := make(map[string]int)
	for _, w := range magazine {
		magMap[w]++
	}

	// iterate through the words in the note decriment the available words when the key is found
	// otherwise if key is not found return false
	for _, w := range note {
		if magMap[w] == 0 {
			return false
		} else {
			magMap[w]--
		}
	}

	return true
}
