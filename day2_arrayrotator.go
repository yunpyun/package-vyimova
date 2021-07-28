package main

import "fmt"

func ArrayRotate(arr []int, count int) []int {
	halfArray1 := arr[:count+1]
	halfArray2 := arr[count+1:]
	var fullArray []int
	fullArray = append(halfArray2, halfArray1...)
	return fullArray
}

func main() {
	arr := []int{3, 8, 9, 7, 6}
	count := 2
	if count < len(arr) {
		res := ArrayRotate(arr, count) // ответ: [7, 6, 3, 8, 9]
		fmt.Println(res)
	} else {
		fmt.Println("Значение count должно быть меньше длины массива!")
	}
}
