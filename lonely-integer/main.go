package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func lonelyinteger(a []int32) int32 {
	tempMap := make(map[int32]int)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, val := range a {
		if _, ok := tempMap[val] ; !ok {
			tempMap[val] = tempMap[val] + 1
		}
	}
	for _, val := range a {
		if _, ok := tempMap[val] ; ok {
			tempMap[val] = tempMap[val] + 1
		}
	}
	var uniq []int32
	for i, val := range tempMap{
		fmt.Println(i, val)
		if val == 2 {
			uniq = append(uniq, i)
		}
	}
	return uniq[0]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := lonelyinteger(a)

	fmt.Fprintf(writer, "%d\n", result)

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
		//panic(err)
	}
}
