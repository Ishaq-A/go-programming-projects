package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.co.uk",
		"Amazon Web Services": "https://aws.com",
	}
	
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	
	websites["LinkedIn"] = "https://linkedin.com"
	
	delete(websites, "Google")
}

