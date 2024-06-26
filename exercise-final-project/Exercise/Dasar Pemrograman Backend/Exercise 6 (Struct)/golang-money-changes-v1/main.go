package main

type Product struct {
	Name  string
	Price int
	Tax   int
}


func MoneyChanges(amount int, products []Product) []int {
	total := 0
	for _, price := range products {
		total += price.Price + price.Tax
	}
	amount -= total

	money := [10]int{-1, 1000, 500, 200, 100, 50, 20, 10, 5, 1}
	change := []int{}

	for i:= 1; i < 10; i++ {
		if amount - money[i] >= 0 {
			change = append(change, money[i])
			amount -= money[i]

			i--
			if money[i] == 0 {
                break
            }
		}
	}
	return change
}
