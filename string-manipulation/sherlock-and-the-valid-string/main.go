package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	if isValidBoolean(s) {
		return "YES"
	}
	return "NO"
}

func isValidBoolean(s string) bool {
	if len(s) < 2 {
		return true
	}
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	countOffBy := 0
	countOffByOne := 0
	initialCount := -1
	initialCountOffByOne := -1
	countOfOnes := 0
	var onesKey rune
	for k, v := range m {
		if v == 1 {
			countOfOnes++
			onesKey = k
		}
		if initialCount < 0 {
			initialCount = v
			initialCountOffByOne = initialCount - 1
			continue
		}
		if initialCount != v {
			offBy := int(math.Abs(float64(initialCount - v)))
			countOffBy += offBy
		}
		if initialCountOffByOne != v {
			offBy := int(math.Abs(float64(initialCountOffByOne - v)))
			countOffByOne += offBy
		}
	}
	pass := (countOffBy == 0 || countOffBy == 1) || (countOffByOne == 0)
	if !pass {
		if countOfOnes == 1 {
			delete(m, onesKey)
			initialCount := -1
			for _, v := range m {
				if initialCount < 0 {
					initialCount = v
					continue
				}
				if initialCount != v {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
