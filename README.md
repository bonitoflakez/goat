# GOAT [Go API Tracer]

I made this utility to simplify API testing. It streamlines the process of examining API requests, responses, and headers, making it easier to analyze and debug API interactions.

## Usage

```txt
>  .\goat.exe -h
Usage: goat -url <URL> -method <METHOD> [OPTIONS]

Options:
  -url <URL>            URL of the API route
  -method <METHOD>      HTTP method of the request [GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS]
  -body <BODY>          Request body for POST, PUT, PATCH requests
  -header <HEADER>      Request headers in the format 'key1:value1,key2:value2'
  -help, -h             Show help message
```

## To-Do

- [x] Implement GET request functionality
- [x] Implement POST request functionality
- [x] Add functionality for PUT requests
- [x] Add functionality for DELETE requests
- [x] Incorporate PATCH request functionality
- [x] Support HEAD requests
- [x] Enable OPTIONS request functionality
- [x] Support for handling request headers (including content type, accept headers, custom headers, etc.)
- [ ] Implement error handling and reporting
