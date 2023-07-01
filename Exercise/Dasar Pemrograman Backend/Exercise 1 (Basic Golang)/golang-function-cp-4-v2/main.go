package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var similar string
	for _, word := range data {
        if strings.Contains(word, input) {
			similar += word + ","
		}
    }
	return similar[:len(similar)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
