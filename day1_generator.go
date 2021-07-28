package main

import "fmt"

func SolutionSquareGenerator(start int, n int) []int {
	var result []int
	result = make([]int, n)
	for x := start; x < start+n; x++ {
		result[x-start] = x * x
	}
	return result
}

func main() {
	start := 6
	n := 18
	generator := SolutionSquareGenerator(start, n)
	fmt.Println(generator)
}
