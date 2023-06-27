package main

import (
	"fmt"
	"net/http"
	"time"
)

type TracingTransport struct {
	Transport http.RoundTripper
}

func (t *TracingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	fmt.Printf("\n[!] Tracing\n\n")
	fmt.Printf("[+] URL -> %s\n", req.URL.String())
	fmt.Printf("[+] Method -> %s\n", req.Method)
	logHeaders("\n[!] Headers ", req.Header)

	resp, err := t.transport().RoundTrip(req)

	duration := time.Since(startTime)

	if err != nil {
		fmt.Println("[!] Error:", err.Error())
	} else {
		fmt.Printf("[+] Response status code -> %d\n", resp.StatusCode)
		logHeaders("\n[!] Response Headers\n", resp.Header)
	}

	fmt.Printf("[+] Response Time -> %s\n", duration)
	return resp, err
}

func (t *TracingTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
