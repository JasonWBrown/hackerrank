package main

import "fmt"

func main() {}

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	// Write your code here
	bribes, err := minimumBribesTestable(q)
	if err != nil {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribes)
	}
}

func minimumBribesTestable(q []int32) (int, error) {
	numberOfBribes := 0
	for i := len(q); i > 0; i-- {
		if i-1 >= 0 && i == int(q[i-1]) {
			continue
		} else if i-2 >= 0 && i == int(q[i-2]) {
			temp := q[i-2]
			q[i-2] = q[i-1]
			q[i-1] = temp
			numberOfBribes++
			continue
		} else if i-3 >= 0 && i == int(q[i-3]) {
			temp := q[i-3]
			q[i-3] = q[i-2]
			q[i-2] = temp

			temp = q[i-2]
			q[i-2] = q[i-1]
			q[i-1] = temp
			numberOfBribes += 2
			continue
		} else {
			return 0, fmt.Errorf("too chaotic")
		}
	}
	return numberOfBribes, nil
}
