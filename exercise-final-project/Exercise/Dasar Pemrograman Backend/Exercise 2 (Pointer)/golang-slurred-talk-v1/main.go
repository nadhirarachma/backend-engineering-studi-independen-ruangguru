package main

import "fmt"

func SlurredTalk(words *string) {
	length := len(*words)
	for i:= 0; i < length; i++ {
		if string((*words)[i]) == "S" || string((*words)[i]) == "R" || string((*words)[i]) == "Z" {
			*words += "L"
		} else if string((*words)[i]) == "s" || string((*words)[i]) == "r" || string((*words)[i]) == "z" {
			*words += "l"
		} else {
			*words += string((*words)[i])
		}
	}
	*words = (*words)[length:]
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = ""
	SlurredTalk(&words)
	fmt.Println(words)
}
