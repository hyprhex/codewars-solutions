package main

import (
	"fmt"
	"math"
)

func main() {
	n := []int{0, 1, 2, 4}
	for _, v := range n {
		fmt.Println(powersOfTwo(v))
	}
}

func powersOfTwo(n int) []uint64 {
	result := []uint64{}
	for i := 0; i <= n; i++ {
		x := math.Pow(2, float64(i))
		result = append(result, uint64(x))
	}

	return result
}
