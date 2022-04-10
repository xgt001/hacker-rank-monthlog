package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int64) {
	var minSum, maxSum int64 = math.MaxInt64, math.MinInt64
	for i := 0; i < len(arr); i++ {
		operated := pop(arr, i)
		for k := 0; k < len(operated); k++ {
			t := sum(operated)
			if t > maxSum {
				maxSum = t
			}
			if t < minSum {
				minSum = t
			}
		}
	}
	fmt.Printf("%d %d", minSum, maxSum)
}

func pop(s []int64, index int) []int64 {
	var poppedSlice []int64
	for i := 0; i < len(s); i++ {
		if i != index {
			poppedSlice = append(poppedSlice, s[i])
		}
	}
	s = poppedSlice
	return s
}

func sum(arr []int64) int64 {
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	//fmt.Println(sum)
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
