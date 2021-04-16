/****************************************/
// Brute Force Solution : O(n^2) 
// .. using for * for
// Kadane's Algorithm : O(n)
// K[i] = Max(K[i-1] + arr[i], k[i])
// https://www.hackerrank.com/challenges/max-array-sum/problem
/****************************************/
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	k := make([]int32, len(arr))

	for i := range arr {
		if i == 0 {
			k[i] = arr[i]
		} else if i == 1 {
			k[i] = max(arr[i], arr[i-1])
		} else {
			k[i] = max(k[i-2]+arr[i], k[i-1])
		}
	}

	if k[len(arr)-1] <= 0 {
		return 0
	}

	return k[len(arr)-1]
}

func max (current int32, next int32) int32 {
	maximum := current
	if current > next {
		maximum = current
	} else {
		maximum = next
	}
	return maximum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := maxSubsetSum(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
