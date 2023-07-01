package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	
	if len(transactions) == 0 {
		return nil
	}

	totalIncome := 0
	totalExpense := 0

	record := make(map[string]string)
	dates := []string{}

	for i := 0; i < len(transactions); i++ {
		
		if transactions[i].Type == "income" {
			totalIncome += transactions[i].Amount
		} else {
			totalExpense += transactions[i].Amount
		}

		totalAmount := 0
		if len(record[transactions[i].Date]) != 0 {
			totalAmount, _ = strconv.Atoi(strings.Split(strings.Split(record[transactions[i].Date], ";")[2], "\n")[0])
		}

		totalAmount += (totalIncome - totalExpense)

		if totalAmount > 0 {
			record[transactions[i].Date] = ";income;" + strconv.Itoa(totalAmount) + "\n"
		} else {
			record[transactions[i].Date] = ";expense;" + strconv.Itoa(totalAmount) + "\n"
		}

		totalIncome = 0
		totalExpense = 0	
	}

	file, err := os.Create(path) 
	if err != nil {
		panic(err)
	}

	defer file.Close() 
	
	for key, _ := range record {
		dates = append(dates, key)
	}

	sort.Strings(dates)
	for i:= 0; i < len(dates); i++ {
		total := record[dates[i]]
		
		if string(string(strings.Split(total, ";")[2][0])) == "-" {
			total =	strings.Replace(total, "-", "", -1)
		}

		if i == len(dates)-1 {
			_, err = file.WriteString(dates[i] + "" + total[0:len(total)-1])
		} else {
			_, err = file.WriteString(dates[i] + "" + total)
		}

		if err != nil {
			panic(err)
		}
	}
	return nil
}


func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
