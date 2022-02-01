package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":            "https://google.com",
		"Amazon Web Server": "https://aws.com",
	}
	fmt.Println(websites)

	websites["LinkedIn"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
