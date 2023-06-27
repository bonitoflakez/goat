# GOAT [Go API Tracer]

I made this utility to simplify API testing. It streamlines the process of examining API requests, responses, and headers, making it easier to analyze and debug API interactions.

## Usage

```txt
>  .\goat.exe -help
usage: goat -url <URL> -method <METHOD>

        -url            URL of API route
        -method         method of route [GET, POST]
        -body           body for requests
        -header         request headers in the format 'key1:value1,key2:value2'
        -help / -h      show help message
```

## To-Do

- [x] Implement GET request functionality
- [x] Implement POST request functionality
- [ ] Add functionality for PUT requests
- [ ] Add functionality for DELETE requests
- [ ] Incorporate PATCH request functionality
- [ ] Support HEAD requests
- [ ] Enable OPTIONS request functionality
- [ ] Support for handling request headers (including content type, accept headers, custom headers, etc.)
- [ ] Implement error handling and reporting
- [ ] Provide options for handling request parameters (query parameters, request body, etc.)
