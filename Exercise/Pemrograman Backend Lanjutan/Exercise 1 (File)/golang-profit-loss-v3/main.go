package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func Readfile(path string) ([]string, error) {
	file, err := os.ReadFile(path) 
	if err != nil {
		panic(err)
	}

	if len(file) == 0 {
		return []string{}, nil
	}

	return strings.Split(string(file), "\n"), nil 
}

func CalculateProfitLoss(data []string) string {

	total := 0
	result := ""

	for i := 0; i < len(data); i++ {
		transaction := strings.Split(data[i], ";")
		amount, _ := strconv.Atoi(transaction[2])

		if transaction[1] == "income" {
			total += amount
		} else {
			total -= amount
		}

		if i == len(data) - 1 {
			result = transaction[0] + ";" 
			if total > 0 {
				result += "profit"
			} else {
				result += "loss"
			}
			result += ";" + strings.Replace(strconv.Itoa(total), "-", "", -1) 
		}
	}
	return result 
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
