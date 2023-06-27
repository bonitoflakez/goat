package main

import (
	"fmt"
	"net/http"
)

func logResponse(resp *http.Response, dump []byte) {
	fmt.Printf("\n\n[!] Response\n\n")
	fmt.Printf("[+] Status Code -> %d\n", resp.StatusCode)

	if len(resp.Header) == 0 {
		fmt.Println("[+] Headers -> NULL")
	} else {
		logHeaders("\n[!] Response Headers\n", resp.Header)
	}

	fmt.Printf("\n[!] Response Body \n\n%s", string(dump))
}
