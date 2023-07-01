package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	var date string

	if day < 10 {
		date += "0" 
	} 
	date += strconv.Itoa(day)

	switch {
		case month == 1:
			date += "-January-"
		case month == 2:
			date += "-February-"
        case month == 3:
			date += "-March-"
		case month == 4:
			date += "-April-"
        case month == 5:
			date += "-May-"
		case month == 6:
			date += "-June-"
        case month == 7:
			date += "-July-"
		case month == 8:
			date += "-August-"
        case month == 9:
			date += "-September-"
		case month == 10:
			date += "-October-"
        case month == 11:
			date += "-November-"
		default:
			date += "-December-"
	}

	date += strconv.Itoa(year)
	return date
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
