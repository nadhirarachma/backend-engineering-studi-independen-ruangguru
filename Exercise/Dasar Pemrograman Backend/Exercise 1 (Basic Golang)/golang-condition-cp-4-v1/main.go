package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {

	totalTicket := VIP + regular + student
	var totalPrice float32 = float32((VIP * 30) + (regular * 20) + (student * 10))

	if totalPrice >= 100 {
		if day % 2 != 0 {
			if totalTicket < 5 {
				totalPrice -= totalPrice * 0.15
			} else {
				totalPrice -= totalPrice * 0.25
			}
		} else {
			if totalTicket < 5 {
				totalPrice -= totalPrice * 0.10
			} else {
				totalPrice -= totalPrice * 0.20
			}
		}
	} 

	return totalPrice

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
