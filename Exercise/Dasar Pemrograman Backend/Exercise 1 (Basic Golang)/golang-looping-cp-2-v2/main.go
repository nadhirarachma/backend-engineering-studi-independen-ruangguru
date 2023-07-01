package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var reversed string

	for i:= len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])

		if i > 0 {
			if string(str[i]) != " " && string(str[i-1]) != " " {
				reversed += "_"
			}
		}
	}

	return reversed
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
