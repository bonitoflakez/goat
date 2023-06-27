package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func logRequest(req *http.Request) {
	fmt.Printf("\n[!] Request\n\n")
	fmt.Printf("[+] URL -> %s\n", req.URL.String())
	fmt.Printf("[+] Method -> %s\n", req.Method)

	if len(req.Header) == 0 {
		fmt.Printf("[!] Headers -> NULL")
	} else {
		logHeaders("\n[!] Headers\n", req.Header)
	}

	if req.Body != nil {
		defer req.Body.Close()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println("[!] Error reading request body: ", err)
		} else {
			if len(body) == 0 {
				fmt.Printf("[!] Request Body -> NULL")
			} else {
				fmt.Printf("\n[!] Request Body \n%s", string(body))
			}
		}
	}
}
