package main

import (
	"fmt"
	"strings"
)

func ReverseWord(str string) string {
	var word string
	var reversed string
	var reversedWord string

	for i:= 0; i < len(str); i++ {
		if string(str[i]) != " " {
			word += string(str[i])
		} 

		if string(str[i]) == " " || i == len(str)-1 {
			for j:= len(word) - 1; j >= 0; j-- {
				if strings.ToUpper(string(word[j])) == string(word[j]) {
					reversed += strings.ToLower(string(word[j]))
					reversedWord += strings.ToUpper(string(reversed[0])) + reversed[1:]
					reversed = ""
					continue
				} 
                reversed += string(word[j])
			}
			word = ""
			reversedWord += reversed + " "
			reversed = ""
		}
	}

	return reversedWord[:len(reversedWord)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
