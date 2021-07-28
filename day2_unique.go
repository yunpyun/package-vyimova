package main

import "fmt"

func GetUniqCount(arr []int) int {
	var result int
	var found bool

	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr); y++ {
			if arr[y] == arr[x] && x != y {
				found = true
				break
			}
		}
		if !found {
			result++
		}
		found = false
	}

	return result
}

func main() {
	arr := []int{3, 8, 9, 7, 3, 5, 8, 9, 6}
	res := GetUniqCount(arr)
	fmt.Println(res) // ответ: 3
}
