package client

import (
    "net/http"
    "net/url"
    "github.com/google/go-querystring/query"
)

type Client struct {
    httpClient *http.Client
}

func (c *Client) List(baseURL string, opts *ListOptions) ([]interface{}, *Response, error) {
    u, err := url.Parse(baseURL)
    if err != nil {
        return nil, nil, err
    }
    
    qs, err := query.Values(opts)
    if err != nil {
        return nil, nil, err
    }
    
    // Remove per_page if it's zero
    if opts.PerPage == 0 {
        qs.Del("per_page")
    }
    
    u.RawQuery = qs.Encode()
    
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, nil, err
    }
    
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, nil, err
    }
    defer resp.Body.Close()
    
    pResp, err := parsePagination(resp)
    if err != nil {
        return nil, nil, err
    }
    
    // Parse items from response body
    // This is a placeholder; replace with actual parsing logic
    items := make([]interface{}, 0)
    
    return items, pResp, nil
}

func addOptions(baseURL string, opts *ListOptions) (string, error) {
    u, err := url.Parse(baseURL)
    if err != nil {
        return "", err
    }
    
    qs, err := query.Values(opts)
    if err != nil {
        return "", err
    }
    
    // Remove per_page if it's zero
    if opts.PerPage == 0 {
        qs.Del("per_page")
    }
    
    u.RawQuery = qs.Encode()
    
    return u.String(), nil
}
