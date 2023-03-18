package main

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	highest := int32(-65) // has to be less than -9*6
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum1 := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			sum2 := arr[i+1][j+1]
			sum3 := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			total := sum1 + sum2 + sum3
			if total > highest {
				highest = total
			}
		}
	}

	return highest
}

func main() {}
