package main

func CountProfit(data [][][2]int) []int {
	totalProfit := []int{}
	if len(data) > 0 {
		totalProfit = make([]int, len(data[0]))
	} 
	
	for i := 0; i < len(data); i++ {
        for j := 0; j < len(data[i]); j++ {
			sum := data[i][j][0] - data[i][j][1]
			totalProfit[j] += sum
		}
    }
	return totalProfit
}

