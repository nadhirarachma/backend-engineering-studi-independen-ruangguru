package main

import (
	"fmt"
)

func FindShortestName(names string) string {
	var shortestName string
	var name string
	var shortest int = 100000
	var length int
	var sum int

	for i:= 0; i < len(names); i++ {

		if string(names[i]) == " " || string(names[i]) == "," || string(names[i]) == ";" {
			length = sum
		    
			if length < shortest {
				shortest = length
				shortestName = name
			} else if length == shortest && name < shortestName {
				shortestName = name
			}
			name = ""
			sum = 0
			continue
		}
		name += string(names[i])
		sum++

		if i == len(names)-1 && sum < shortest {
			shortest = sum
			shortestName = name
		} else if i == len(names)-1 && sum == shortest && name < shortestName {
			shortestName = name
		}
	}
	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tio;Tia"))                         // "Tia"
}
