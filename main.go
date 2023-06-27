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
	fmt.Println("Usage: goat -url <URL> -method <METHOD> [OPTIONS]")
	fmt.Println("")

	fmt.Println("Options:")
	fmt.Println("  -url <URL>\t\tURL of the API route")
	fmt.Println("  -method <METHOD>\tHTTP method of the request [GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS]")
	fmt.Println("  -body <BODY>\t\tRequest body for POST, PUT, PATCH requests")
	fmt.Println("  -header <HEADER>\tRequest headers in the format 'key1:value1,key2:value2'")
	fmt.Println("  -help, -h\t\tShow help message")
}
