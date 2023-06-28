package main

import (
	"fmt"
	"net/http"
	"strings"
)

func printHeaders(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			fmt.Printf("[+] %s: %s\n", key, value)
		}
	}
}

func parseHeaders(headerString string) http.Header {
	headers := http.Header{}
	headerList := strings.Split(headerString, ",")
	for _, header := range headerList {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers.Add(key, value)
		}
	}
	return headers
}
