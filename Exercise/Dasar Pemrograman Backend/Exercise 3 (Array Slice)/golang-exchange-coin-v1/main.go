package main

func ExchangeCoin(amount int) []int {
	if amount == 0 {
		return []int{}
	} 

	coin := [10]int{-1, 1000, 500, 200, 100, 50, 20, 10, 5, 1}
	exchange := []int{}
	

	for i:= 1; i < 10; i++ {
		if amount - coin[i] >= 0 {
			exchange = append(exchange, coin[i])
			amount -= coin[i]

			i--
			if coin[i] == 0 {
                break
            }
		}
	}
	return exchange 
}
