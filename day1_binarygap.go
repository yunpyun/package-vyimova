package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionBinaryGap(N int) int {
	stringN := strconv.Itoa(N)           // конвертируем int в string
	arrayN := strings.Split(stringN, "") // делим строку на массив из нулей и единиц
	result := 0                          // переменная для выводы итогового числа нулей
	counter := 0                         // счетчик для подсчета значений
	for x := 0; x < len(stringN); x++ {
		if arrayN[x] == "0" {
			counter++
		}
		if arrayN[x] == "1" {
			if counter > result {
				result = counter
			}
			counter = 0
		}
	}
	if counter > result {
		result = counter
	}
	return result
}

func main() {
	N := 1000111000010 // 4546 в десятичном представлнии
	binaryGap := SolutionBinaryGap(N)
	fmt.Println(binaryGap) // выведет количество нулей самой длинной цепочки, ответ - 4
}
