package main

import (
	"fmt"
	"strings"
	"strconv"
)

func ChangeOutput(data []string) map[string][]string {
	output := make(map[string][]string)

	for i:= 0; i < len(data); i++ {
		val := strings.Split(data[i], "-")
		index, _ := strconv.Atoi(val[1])

		if val[2] == "last" {
			pair := output[val[0]][index] + " " + val[3]
			output[val[0]][index] = pair
		} else {
			output[val[0]] = append(output[val[0]], val[3])
		}
	}
	return output
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
