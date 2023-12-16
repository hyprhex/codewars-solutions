package main

import "fmt"

func main() {
	m := map[int]int {1:5, 2:5, 3:5, 50:5, 100:5}
	for k, v := range m {
		fmt.Println(countBy(k, v))
	}
}

func countBy(x, n int) []int {

	result := []int{}

	if x > 0 && n > 0 {
		for i := 1; i <= n; i++ {
			result = append(result, i*x)
		}
	}

	return result

}
