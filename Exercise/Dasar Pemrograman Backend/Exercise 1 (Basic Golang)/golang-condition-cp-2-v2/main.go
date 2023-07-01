package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var heightFloat float64 = float64(height)
	var bmi float64
	
	if gender == "laki-laki" {
		bmi = (heightFloat - 100) - ((heightFloat - 100) * 0.1)
	} else {
		bmi = (heightFloat - 100) - ((heightFloat - 100) * 0.15)
	}
	return bmi 
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
