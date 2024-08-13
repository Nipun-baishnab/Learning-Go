package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"https://hackernoon.com/",
		"https://github.com/",
		"https://apple.com/",
		"https://google.com/",
		"https://youtube.com/",
		"https://www.udemy.com/",
		"https://netflix.com/",
		"https://www.coursera.org/",
		"https://facebook.com/",
		"https://microsoft.com",
		"https://wikipedia.org",
		"https://educative.io",
		"https://acloudguru.com",
		"https://www.codeforces.com",
	}

	for _, website := range websites {
		checkResource(website)
	}

}
func checkResource(website string) {
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")

	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}
}
