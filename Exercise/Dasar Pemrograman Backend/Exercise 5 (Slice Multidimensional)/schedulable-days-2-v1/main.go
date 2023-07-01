package main

import "sort"

func SchedulableDays(villager [][]int) []int {
	date := []int{}

	if len(villager) == 1 {
		return villager[0]
	}

	for m:= 0; m < len(villager); m++ {
		for n:= 0; n < len(villager[m]); n++ {
			date = append(date, villager[m][n])
		}
	}
	sort.Ints(date)

	schedulable := []int{}
	for i:= 0 ; i < len(date)-1; i++ {
		if date[i] == date[i+1] {
            schedulable = append(schedulable, date[i])
        }
	}
	return schedulable
}

