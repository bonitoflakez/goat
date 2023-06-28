package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func logRequest(req *http.Request) {
	switch outputFormat {
	case "simple":
		fmt.Printf("[+] URL: %s\n", req.URL.String())
		fmt.Printf("[+] Method: %s\n", req.Method)
		fmt.Println("[+]")
	case "detailed":
		fmt.Println("[!] Request")
		fmt.Printf("[+] URL: %s\n", req.URL.String())
		fmt.Printf("[+] Method: %s\n", req.Method)
		fmt.Println("[!] Headers:")
		printHeaders(req.Header)
		fmt.Println()
	case "full":
		fmt.Println("[!] Request")
		fmt.Printf("[+] URL: %s\n", req.URL.String())
		fmt.Printf("[+] Method: %s\n", req.Method)
		fmt.Println("[!] Headers:")
		printHeaders(req.Header)
		fmt.Println()
		fmt.Println("[!] Full Request:")
		fmt.Println(stringRequest(req))
		fmt.Println()
	default:
		fmt.Printf("Invalid output format: %s\n", outputFormat)
		os.Exit(1)
	}
}

func stringRequest(req *http.Request) string {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}
	return string(dump)
}
