package main

import "fmt"

func TicketPlayground(height, age int) int {
	var price int

	switch {
		case age < 5:
			price = -1
		case age > 12:
			price = 100000
		case age == 12 || height > 160:
			price = 60000
		case (age >= 10 && age <= 11) || height > 150:
			price = 40000
		case (age >= 8 && age <= 9) || height > 135:
			price =  25000
		case (age >= 5 && age <= 7) || height > 120:
			price = 15000
	}
	return price
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
