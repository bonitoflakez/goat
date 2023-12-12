# GOAT [Go API Tracer]

I made this utility to simplify API testing. It streamlines the process of examining API requests, responses, and headers, making it easier to analyze and debug API interactions.

### Usage

```txt
>  .\goat.exe -help
usage: goat -url <URL> -method <METHOD> [OPTIONS]

        -url            URL of API route
        -method         method of route [GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS]
        -body           body for requests
        -header         request headers in the format 'key1:value1,key2:value2'
        -help / -h      show help message
        -output         output format [simple, detailed, full]
```

### Things GOAT does

- Trace GET, POST, PUT, DELETE, PATCH, OPTIONS requests
- Supports HEAD requests
- Can handle request headers (including content type, accept headers, custom headers, etc.)
