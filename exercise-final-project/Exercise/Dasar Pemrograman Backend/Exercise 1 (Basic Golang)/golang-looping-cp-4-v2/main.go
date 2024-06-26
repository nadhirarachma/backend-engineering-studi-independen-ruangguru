package main

import "fmt"

func EmailInfo(email string) string {

	var domain string
	var tld string
	var domainIndex int
	var tldIndex int

	for i:= 0; i < len(email); i++ {

		if string(email[i]) == "@" {
			domainIndex = i
		}

		if string(email[i]) == "." {
			tldIndex = i
			break
		}
	}

	domain = string(email[domainIndex + 1:tldIndex])
	tld = string(email[tldIndex + 1:])
	return "Domain: " + domain + " dan " + "TLD: " + tld
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
