package main

import "strconv"

func ReverseData(arr [5]int) [5]int {
	var num string

	for i:= 0; i < 5; i++ {
		for j:= len(strconv.Itoa(arr[i]))-1; j >= 0; j-- {
			num += string(strconv.Itoa(arr[i])[j])
		}
		arr[i], _ = strconv.Atoi(num)
		num = ""
	}

	reversed:= [5]int{}
	for m, n:= 0, 4; m < 5; m++ {
		reversed[m] = arr[n]
		n--
	}

	return reversed
}