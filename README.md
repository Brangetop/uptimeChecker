# Concurrent Website Checker

A lightweight Go command-line tool that concurrently checks the availability and response times of multiple websites.

## Features

- **Concurrent Execution**: Uses Go goroutines and `sync.WaitGroup` to check multiple URLs simultaneously.
- **Performance Tracking**: Measures and prints the exact response duration for successful connections.
- **Error Handling**: Detects and reports connection timeouts, invalid URLs, and non-2xx HTTP status codes.
- **Custom Timeout**: Configured with a 5-second request timeout to prevent hanging.

## How It Works

1. The program iterates through a predefined list of website URLs.
2. A new goroutine is spawned for each URL to perform an HTTP `HEAD` request.
3. The `sync.WaitGroup` ensures the main program waits for all concurrent checks to complete before exiting.

## Prerequisites

- Go 1.16 or higher

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com
   cd your-repo-name
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

## Example Output

```text
https://github.com OK Duration:  181.672171ms
https://wikipedia.org/ FAIL StatusCode:  403
https://go.dev/ OK Duration:  296.875704ms
https://google.com/ OK Duration:  388.678574ms
https://youtube.com/ ERROR: Head "https://youtube.com/": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```
