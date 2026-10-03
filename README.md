# Client Library

A Go client library for interacting with the API.

## Installation
```sh
go get github.com/yourusername/client
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    "github.com/yourusername/client"
)

func main() {
    // Create a new client.
    c := client.NewClient(&http.Client{}, "https://api.example.com")

    // List items with pagination.
    items, pResp, err := c.List("/items", &client.ListOptions{Page: 1, PerPage: 0})
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the items.
    for _, item := range items {
        fmt.Println(item)
    }

    // Print the pagination information.
    fmt.Println("Next Page:", pResp.NextPage())
    fmt.Println("Prev Page:", pResp.PrevPage())
    fmt.Println("First Page:", pResp.FirstPage())
    fmt.Println("Last Page:", pResp.LastPage())
}
```

## Pagination
The client library supports pagination. To use pagination, specify the `Page` and `PerPage` parameters in the `ListOptions` struct. If `PerPage` is set to `0`, the `per_page` parameter will be omitted from the request URL.

## License
MIT
