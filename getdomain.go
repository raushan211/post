package main

import (
	"fmt"
	"net/url"

	//"net/url"
	"strings"
)

func getdomain(ur string) (string, error) {
	domain := ""
	url, err := url.Parse(ur)
	if err != nil {
		return domain, err
	}
	parts := strings.Split(url.Hostname(), ".")
	domain = parts[len(parts)-2] + "." + parts[len(parts)-1]
	fmt.Println(domain)
	return domain, err
}
