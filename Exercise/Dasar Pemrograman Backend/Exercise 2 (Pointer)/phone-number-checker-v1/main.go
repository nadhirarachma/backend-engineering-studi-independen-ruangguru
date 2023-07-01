package main

import (
	"fmt"
	"strconv"
)

func PhoneNumberChecker(number string, result *string) {
	var code int

	if string(number[:2]) == "08" && len(number) >= 10 {
		code, _ = strconv.Atoi(string(number[2:4]))
	} else if string(number[:2]) == "62" && len(number) >= 11 {
		code, _ = strconv.Atoi(string(number[3:5]))
	} else {
		*result = "invalid"
	}

	switch {
		case code >= 11 && code <= 15:
			*result = "Telkomsel"
		case code >= 16 && code <= 19:
			*result = "Indosat"
		case code >= 21 && code <= 23:
			*result = "XL"
		case code >= 27 && code <= 29:
			*result = "Tri"
		case code >= 52 && code <= 53:
			*result = "AS"
		case code >= 81 && code <= 88:
			*result = "Smartfren"
		default:
			*result = "invalid"
	}
	
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "08222"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
