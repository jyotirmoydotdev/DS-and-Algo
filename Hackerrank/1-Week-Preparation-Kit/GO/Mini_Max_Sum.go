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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var max int32 = -2147483648
	var min int32 = 2147483647
	for i := 0; i < len(arr); i++ {
		if max <= arr[i] {
			max = arr[i]
		}
		if min >= arr[i] {
			min = arr[i]
		}
	}

	var maxSum int64
	var minSum int64
	if max == min {
		maxSum = int64(max) * 4
		minSum = int64(min) * 4
	} else {
		for i := 0; i < len(arr); i++ {
			if arr[i] != max {
				minSum += int64(arr[i])
			}
			if arr[i] != min {
				maxSum += int64(arr[i])
			}
		}
	}
	fmt.Println(minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
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
