package main

import (
	"fmt"
	"log"
	"net/url"

	//"net/url"
	"strings"
)

func getdomain(ur string) string {
	url, err := url.Parse(ur)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(url.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Println(domain)
	return domain
}
