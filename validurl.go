package main

import (
	"fmt"
	"net/url"
)

func validUrl(ur string) string {
	url, err := url.ParseRequestURI(ur)
	if err != nil {
		panic(err)
	}
	u, err := url.Parse(ur)
	if err != nil || u.Scheme == "" || u.Host == "" {
		fmt.Println("Not ok")
	}

	fmt.Println("All ok")
	return ur
}
