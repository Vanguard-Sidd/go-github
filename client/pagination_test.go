package client

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestAddOptions(t *testing.T) {
    baseURL := "https://api.example.com/items"

    // Test that PerPage is omitted when zero.
    opts := &ListOptions{Page: 1, PerPage: 0}
    url, err := addOptions(baseURL, opts)
    if err != nil {
        t.Fatal(err)
    }
    if url != "https://api.example.com/items?page=1" {
        t.Errorf("Expected URL to be https://api.example.com/items?page=1, got %s", url)
    }

    // Test that PerPage is included when non-zero.
    opts = &ListOptions{Page: 1, PerPage: 10}
    url, err = addOptions(baseURL, opts)
    if err != nil {
        t.Fatal(err)
    }
    if url != "https://api.example.com/items?page=1&per_page=10" {
        t.Errorf("Expected URL to be https://api.example.com/items?page=1&per_page=10, got %s", url)
    }
}

func TestParsePagination(t *testing.T) {
    // Test that parsePagination correctly parses the Link header.
    linkHeader := "<https://api.example.com/items?page=2>; rel=\"next\", <https://api.example.com/items?page=1>; rel=\"prev\", <https://api.example.com/items?page=1>; rel=\"first\", <https://api.example.com/items?page=5>; rel=\"last\"
    resp := &http.Response{
        Header: http.Header{
            "Link": []string{linkHeader},
        },
    }
    pResp, err := parsePagination(resp)
    if err != nil {
        t.Fatal(err)
    }
    if pResp.NextPage != 2 {
        t.Errorf("Expected NextPage to be 2, got %d", pResp.NextPage)
    }
    if pResp.PrevPage != 1 {
        t.Errorf("Expected PrevPage to be 1, got %d", pResp.PrevPage)
    }
    if pResp.FirstPage != 1 {
        t.Errorf("Expected FirstPage to be 1, got %d", pResp.FirstPage)
    }
    if pResp.LastPage != 5 {
        t.Errorf("Expected LastPage to be 5, got %d", pResp.LastPage)
    }

    // Test that parsePagination correctly handles a missing page parameter.
    linkHeader = "<https://api.example.com/items>; rel=\"next\", <https://api.example.com/items>; rel=\"prev\", <https://api.example.com/items>; rel=\"first\", <https://api.example.com/items>; rel=\"last\"
    resp = &http.Response{
        Header: http.Header{
            "Link": []string{linkHeader},
        },
    }
    pResp, err = parsePagination(resp)
    if err != nil {
        t.Fatal(err)
    }
    if pResp.NextPage != 0 {
        t.Errorf("Expected NextPage to be 0, got %d", pResp.NextPage)
    }
    if pResp.PrevPage != 0 {
        t.Errorf("Expected PrevPage to be 0, got %d", pResp.PrevPage)
    }
    if pResp.FirstPage != 0 {
        t.Errorf("Expected FirstPage to be 0, got %d", pResp.FirstPage)
    }
    if pResp.LastPage != 0 {
        t.Errorf("Expected LastPage to be 0, got %d", pResp.LastPage)
    }
}

func TestResponseHelpers(t *testing.T) {
    // Test that the response helpers correctly return the page numbers.
    pResp := &Response{
        NextPage:  2,
        PrevPage:  1,
        FirstPage: 1,
        LastPage:  5,
    }
    if pResp.NextPage() != 2 {
        t.Errorf("Expected NextPage to be 2, got %d", pResp.NextPage())
    }
    if pResp.PrevPage() != 1 {
        t.Errorf("Expected PrevPage to be 1, got %d", pResp.PrevPage())
    }
    if pResp.FirstPage() != 1 {
        t.Errorf("Expected FirstPage to be 1, got %d", pResp.FirstPage())
    }
    if pResp.LastPage() != 5 {
        t.Errorf("Expected LastPage to be 5, got %d", pResp.LastPage())
    }
}
