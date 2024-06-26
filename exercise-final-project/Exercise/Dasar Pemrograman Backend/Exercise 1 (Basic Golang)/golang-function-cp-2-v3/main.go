package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var vowel int
	var consonant int
	var check bool

	for i:= 0; i < len(str); i++ {
		if string(str[i]) == "a" || string(str[i]) == "i" || string(str[i]) == "u" || string(str[i]) == "e" || string(str[i]) == "o" || 
		string(str[i]) == "A" || string(str[i]) == "I" || string(str[i]) == "U" || string(str[i]) == "E" || string(str[i]) == "O" {
            vowel++
        } else if str[i] <= 64 || str[i] >= 91 && str[i] <= 96 || str[i] >= 123 {
			continue
		} else {
			consonant++
		}

	}

	if vowel == 0 || consonant == 0 {
		check = true
	}
	return vowel, consonant, check
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
