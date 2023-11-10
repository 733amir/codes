package main

import "fmt"

func main() {

    tcs := [][]int{
        {17,18,5,4,6,1},
    }

    for _, tc := range tcs {
        fmt.Printf("%v => %v\n", tc, replaceElements(tc))
    }
}

func replaceElements(arr []int) []int {
    mx := -1

    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] > mx {
            arr[i], mx = mx, arr[i]
        } else {
            arr[i] = mx
        }
    }

    return arr
}
