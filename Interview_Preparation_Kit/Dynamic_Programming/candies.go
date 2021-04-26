/****************************************/
// O(n) 
// Left To Right / Right to Left
// Maximum value is the result
// https://www.hackerrank.com/challenges/candies/problem
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

// Complete the candies function below.
func candies(n int32, arr []int32) int64 {
    var result int64
    LtoR := make([]int32, len(arr))
    RtoL := make([]int32, len(arr))

    LtoR[0] = 1
    RtoL[len(arr)-1] = 1
    
    for i := 1; i < len(arr); i++ {
        if arr[i-1] < arr[i] {
            LtoR[i] = LtoR[i-1] + 1
        } else {
            LtoR[i] = 1
        }
    }
    
    for i := len(arr)-2; i >= 0; i-- {
        if arr[i] <= arr[i+1] {
            RtoL[i] = 1
        } else {
            RtoL[i] = RtoL[i+1] + 1
        }
    }
    result = compare(LtoR, RtoL)
    
    return result
}

func compare(ltor []int32, rtol []int32) int64 {
    sum := int64(0)
    for i := range ltor {
        if ltor[i] > rtol[i] {
            sum += int64(ltor[i])
        } else {
            sum += int64(rtol[i])
        }
    }
    return sum
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

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := candies(n, arr)

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
        panic(err)
    }
}