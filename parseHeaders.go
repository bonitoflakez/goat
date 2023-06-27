package main

import "strings"

func parseHeaders(headerStr string) map[string][]string {
	headers := make(map[string][]string)
	pairs := strings.Split(headerStr, ",")
	for _, pair := range pairs {
		if kv := strings.Split(pair, ":"); len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			headers[key] = append(headers[key], value)
		}
	}
	return headers
}
