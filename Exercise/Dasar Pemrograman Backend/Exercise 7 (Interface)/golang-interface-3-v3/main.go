package main

import (
	"fmt"
	"reflect"
	"strings"
	"strconv"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	hour := 0
	minute := 0
	
	if reflect.TypeOf(time).String() == "string" && len(time.(string)) == 5 {
		data := strings.Split(time.(string), ":")
		hour, _ = strconv.Atoi(string(data[0]))
		minute, _ = strconv.Atoi(data[1])
	} else if reflect.TypeOf(time).String() == "[]int" && len(time.([]int)) == 2 {
		hour = time.([]int)[0]
		minute = time.([]int)[1]
	} else if reflect.TypeOf(time).String() == "map[string]int" {
		hourMap, hourExist := time.(map[string]int)["hour"]
		minuteMap, minuteExist := time.(map[string]int)["minute"]

		if hourExist && minuteExist {
            hour = hourMap
            minute = minuteMap
        } else {
			return "Invalid input"
		}
	} else if reflect.TypeOf(time).String() == "main.Time" {
		hour = time.(Time).Hour
		minute = time.(Time).Minute
	} else {
		return "Invalid input"
	}

	format := ":" + strconv.Itoa(minute)
	if hour > 0 && hour < 12 {
		format = "0" + strconv.Itoa(hour) + format
		if minute < 10 {
			format += "0"
		}
		return format + " AM"
	} else if hour == 12 {
		return strconv.Itoa(hour) + ":00 PM"
	} else {
		format = strconv.Itoa(hour - 12) + format
		if hour - 12 < 10 {
			format = "0" + format
		}

		if minute < 10 {
			format += "0"
		}
		return format + " PM"
	}
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
