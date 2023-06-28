package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func logResponse(resp *http.Response, dump []byte) {
	switch outputFormat {
	case "simple":
		fmt.Printf("[+] Status Code: %d\n", resp.StatusCode)
		fmt.Printf("[+] Content-Type: %s\n", resp.Header.Get("Content-Type"))
		fmt.Println("[+]")
		fmt.Println("[+] Response Body:")
		printResponseBody(resp.Body)
		fmt.Println()
	case "detailed":
		fmt.Println("[!] Response")
		fmt.Printf("[+] Status Code: %d\n", resp.StatusCode)
		fmt.Println("[!] Response Headers:")
		printHeaders(resp.Header)
		fmt.Println()
		fmt.Println("[+] Response Body:")
		printResponseBody(resp.Body)
		fmt.Println()
	case "full":
		fmt.Println("[!] Response")
		fmt.Printf("[+] Status Code: %d\n", resp.StatusCode)
		fmt.Println("[!] Response Headers:")
		printHeaders(resp.Header)
		fmt.Println()
		fmt.Println("[!] Full Response:")
		fmt.Println(string(dump))
		fmt.Println()
	default:
		fmt.Printf("Invalid output format: %s\n", outputFormat)
		os.Exit(1)
	}
}

func printResponseBody(body io.ReadCloser) {
	io.Copy(os.Stdout, body)
	fmt.Println()
}
