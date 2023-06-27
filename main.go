package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var showHelp bool

func main() {
	routeURL := flag.String("url", "", "URL of API route")
	method := flag.String("method", "", "Method of route")
	requestBody := flag.String("body", "", "Request body for POST request")
	requestHeaders := flag.String("header", "", "Request headers in the format 'key1:value1,key2:value2'")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.BoolVar(&showHelp, "h", false, "Show help message")

	flag.Parse()

	if showHelp {
		help()
		os.Exit(1)
	}

	if flag.NFlag() == 0 {
		fmt.Println("Use -help or -h to see help message")
		os.Exit(1)
	}

	client := &http.Client{
		Transport: &TracingTransport{},
	}

	req, err := http.NewRequest(*method, *routeURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	if *method != "" {
		req.Method = *method
	}

	if *method == "POST" || *method == "PUT" || *method == "PATCH" {
		req.Body = io.NopCloser(strings.NewReader(*requestBody))
	}

	if *requestHeaders != "" {
		headers := parseHeaders(*requestHeaders)
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	logRequest(req)
	logResponse(resp, dump)
}

func help() {
	fmt.Printf("usage: goat -url <URL> -method <METHOD> [OPTIONS]\n\n")
	fmt.Printf("\t-url\t\tURL of API route\n")
	fmt.Printf("\t-method\t\tmethod of route [GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS]\n")
	fmt.Printf("\t-body\t\tbody for requests\n")
	fmt.Printf("\t-header\t\trequest headers in the format 'key1:value1,key2:value2'\n")
	fmt.Printf("\t-help / -h\tshow help message\n")
}
