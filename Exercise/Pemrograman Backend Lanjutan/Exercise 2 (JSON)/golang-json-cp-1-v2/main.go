package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	StudyName string `json:"study_name"`
	StudyCredit int `json:"study_credit"`
	Grade string `json:"grade"`
}

type Report struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Semester int `json:"semester"`
	Studies []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {

	var jsonData, err = ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

	var report Report
	err = json.Unmarshal(jsonData, &report)
    if err != nil {
        panic(err)
    }

	return report, err
}

func GradePoint(report Report) float64 {
	score := 0.0
	credit := 0.0

	if len(report.Studies) == 0 {
		return 0.0
	}

	for _, study := range report.Studies {
		
		if study.Grade == "A" {
			score += 4.0 * float64(study.StudyCredit)
		} else if study.Grade == "AB" {
			score += 3.5 * float64(study.StudyCredit)
		} else if study.Grade == "B" {
			score += 3.0 * float64(study.StudyCredit)
		} else if study.Grade == "BC" {
			score += 2.5 * float64(study.StudyCredit)
		} else if study.Grade == "C" {
			score += 2.0 * float64(study.StudyCredit)
		} else if study.Grade == "CD" {
			score += 1.5 * float64(study.StudyCredit)
		} else if study.Grade == "D" {
			score += 1.0 * float64(study.StudyCredit)
		} else if study.Grade == "DE" {
			score += 0.5 * float64(study.StudyCredit)
		} else {
			score += 0 
		}

		credit += float64(study.StudyCredit)
	}
	
	return score/credit
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
