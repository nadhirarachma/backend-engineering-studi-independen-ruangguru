package main

import(
	"fmt"
	"strings"
	"strconv"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	delivery := make(map[string]float32)
	locationByDay := map[string][]string {
		"senin": {"JKT", "DPK"},
		"selasa": {"JKT", "DPK", "BKS"},
		"rabu": {"JKT", "BDG"},
		"kamis": {"JKT", "BDG", "BKS"},
		"jumat": {"JKT", "BKS"},
		"sabtu": {"JKT", "BDG"},
	}

	for i:= 0; i < len(data); i++ {
		code := strings.Split(data[i], ":")[3]
		name := strings.Split(data[i], ":")[0] + "-" + strings.Split(data[i], ":")[1]
		price, _ := strconv.Atoi(strings.Split(data[i], ":")[2]) 
		totalPrice := float32(price) * 0.05 + float32(price)

		for _, loc := range locationByDay[day] {
			if code == loc {
				if day == "senin" || day == "rabu" || day == "jumat" {
					totalPrice += float32(price) * 0.05 
				}
				delivery[name] = totalPrice
			}
		}
	}

	return delivery
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
