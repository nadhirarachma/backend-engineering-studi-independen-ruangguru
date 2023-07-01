package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var count int

	for i:= 0; i < len(text); i++ {
		if string(text[i]) == "R" || string(text[i]) == "S" || string(text[i]) == "T" || string(text[i]) == "Z" ||
		string(text[i]) == "r" || string(text[i]) == "s" || string(text[i]) == "t" || string(text[i]) == "z" {
            count++
        }
	}
	return count

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
