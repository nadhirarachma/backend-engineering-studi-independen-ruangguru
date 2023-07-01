package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	var str = strconv.Itoa(numbers)
	num1, _ := strconv.Atoi(string(str[0]))
	num2, _ := strconv.Atoi(string(str[1]))

	var biggest int = num1 + num2
	var pair string = string(str[0]) + string(str[1])
	res, _ := strconv.Atoi(pair)

	for i:= 1; i < len(str)-1; i++ {
		num1, _ := strconv.Atoi(string(str[i])) 
		num2, _ := strconv.Atoi(string(str[i+1]))

		sum := num1 + num2

		if sum > biggest {
			biggest = sum
            pair = string(str[i]) + string(str[i+1])
			res, _ = strconv.Atoi(pair)
		}
	}
	return res
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
