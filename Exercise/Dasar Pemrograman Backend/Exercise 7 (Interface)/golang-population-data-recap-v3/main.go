package main

import(
	"strings"
	"strconv"
)

func PopulationData(data []string) []map[string]interface{} {
	dataRecap := make([]map[string]interface{}, len(data))

	if data == nil {
		return dataRecap
	}

	for i:= 0; i < len(data); i++ {
		detail := strings.Split(data[i], ";")
		populationData := make(map[string]interface{})
		populationData["name"] = detail[0]
		populationData["age"], _ = strconv.Atoi(detail[1])
		populationData["address"] = detail[2]

		if detail[3] != "" {
			populationData["height"], _ = strconv.ParseFloat(detail[3], 64)
		} 

		if detail[4]!= "" {
            populationData["isMarried"], _ = strconv.ParseBool(detail[4])
		} 
		dataRecap[i] = populationData
	}
	return dataRecap
}

