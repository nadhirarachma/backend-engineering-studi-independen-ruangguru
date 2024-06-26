package main

import (
	"fmt"
	"errors"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	TLD, IDN_TLD := GetTLD(website.Domain)

	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
	} else if website.Valid == false {
		chErr <- errors.New("domain not valid")
	} else if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
	} else {
		websiteData := &website
		websiteData.TLD = TLD
		websiteData.IDN_TLD = IDN_TLD

		ch <- *websiteData
		chErr <- nil
	}
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	filteredData := []RowData{}
	var err error
	for i:= 0; i < len(data); i++ {
		err = <-errCh
		if err == nil {
			processedTLD := <-ch
			if processedTLD.TLD == TLD {
				filteredData = append(filteredData, processedTLD)
			}
		} else {
			return []RowData{}, err
		}
	}

	return filteredData, err
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
