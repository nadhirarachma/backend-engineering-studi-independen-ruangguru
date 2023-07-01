package main

import (
	"fmt"
	"strings"
	"strconv"
	"a21hc3NpZ25tZW50/internal"
)


func AdvanceCalculator(calculate string) float32 {
	
	if calculate == "" {
		result := internal.NewCalculator(float32(0))
		
		return result.Result()
	} else if len(calculate) == 1 {
		num, _ := strconv.ParseFloat(calculate, 32)
		result := internal.NewCalculator(float32(num))

		return result.Result()
	} else {
		input := strings.Split(calculate, " ")
		num, _ := strconv.ParseFloat(input[0], 32)
		result := internal.NewCalculator(float32(num))

		for i := 2; i < len(input); i+=2 {
			num, _ := strconv.ParseFloat(input[i], 32)
			
			if input[i-1] == "+" {
				result.Add(float32(num))
			} else if input[i-1] == "-" {
				result.Subtract(float32(num))
			} else if input[i-1] == "*" {
				result.Multiply(float32(num))
			} else if input[i-1] == "/" {
				result.Divide(float32(num))
			} 
		}
		return result.Result()
	}
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
