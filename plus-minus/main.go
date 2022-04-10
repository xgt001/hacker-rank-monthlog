package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func plusMinus(arr []int32) {
	var slicePositive []int32
	var sliceNegative []int32
	var sliceZero []int32

	//    approach 1, brute force
	for i := 0; i < len(arr); i++ {
		switch {
		case arr[i] > 0:
			slicePositive = append(slicePositive, arr[i])
		case arr[i] < 0:
			sliceNegative = append(sliceNegative, arr[i])
		default:
			sliceZero = append(sliceZero, arr[i])
		}
	}
	var ratioPositive = float32(len(slicePositive)) / float32(len(arr))
	var ratioNegative = float32(len(sliceNegative)) / float32(len(arr))
	var ratioZero = float32(len(sliceZero)) / float32(len(arr))
	fmt.Printf("%.6f\n", ratioPositive)
	fmt.Printf("%.6f\n", ratioNegative)
	fmt.Printf("%.6f\n", ratioZero)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
