package main

import "sort"

func SchedulableDays(date1 []int, date2 []int) []int {
	date := append(date1, date2...)
	sort.Ints(date)

	schedulable := []int{}
	for i:= 0 ; i < len(date)-1; i++ {
		if date[i] == date[i+1] {
            schedulable = append(schedulable, date[i])
        }
	}
	return schedulable
}

