package main

import (
	"fmt"
	"net/http"
)

func logHeaders(title string, headers http.Header) {
	if len(headers) == 0 {
		fmt.Println("[+] " + title + " -> NULL")
	} else {
		fmt.Println(title)
		for key, values := range headers {
			fmt.Printf("[+] %s : ", key)
			for _, value := range values {
				fmt.Printf("%s", value)
			}
			fmt.Println()
		}
	}
}
