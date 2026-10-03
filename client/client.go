package client

import (
    "encoding/json"
    "net/http"
    "net/url"
    "github.com/google/go-querystring/query"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
//
// Note that `PerPage` is omitted when zero.
type ListOptions struct {
    Page    int `url:"page,omitempty"`
    PerPage int `url:"per_page,omitempty"`
}

// Response wraps an HTTP response with pagination information.
type Response struct {
    *http.Response
    NextPage   int
    PrevPage   int
    FirstPage  int
    LastPage   int
}

// Client represents a client for interacting with the API.
type Client struct {
    httpClient *http.Client
    baseURL    string
}

// NewClient creates a new Client instance.
func NewClient(httpClient *http.Client, baseURL string) *Client {
    return &Client{
        httpClient: httpClient,
        baseURL:    baseURL,
    }
}

// List performs a GET request to the specified endpoint with the given options.
// It returns the parsed items, pagination response, and any error encountered.
func (c *Client) List(endpoint string, opts *ListOptions) ([]string, *Response, error) {
    // Build the URL with the given options.
    url, err := addOptions(c.baseURL+endpoint, opts)
    if err != nil {
        return nil, nil, err
    }

    // Create a new request.
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, nil, err
    }

    // Perform the request.
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, nil, err
    }
    defer resp.Body.Close()

    // Parse the pagination information.
    pResp, err := parsePagination(resp)
    if err != nil {
        return nil, nil, err
    }

    // Parse the response body.
    var items []string
    if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
        return nil, nil, err
    }

    return items, pResp, nil
}

// addOptions adds the pagination options to the base URL.
func addOptions(baseURL string, opts *ListOptions) (string, error) {
    // Parse the base URL.
    u, err := url.Parse(baseURL)
    if err != nil {
        return "", err
    }

    // Encode the options.
    v, err := query.Values(opts)
    if err != nil {
        return "", err
    }

    // Remove the per_page parameter if it is zero.
    if opts.PerPage == 0 {
        v.Del("per_page")
    }

    // Merge the encoded values with the existing query parameters.
    if u.RawQuery != "" {
        existing, err := url.ParseQuery(u.RawQuery)
        if err != nil {
            return "", err
        }
        for key, values := range existing {
            if _, ok := v[key]; !ok {
                v[key] = values
            }
        }
    }

    // Set the new query string.
    u.RawQuery = v.Encode()

    return u.String(), nil
}

// parsePagination parses the pagination information from the response.
func parsePagination(resp *http.Response) (*Response, error) {
    r := &Response{Response: resp}
    linkHeader := resp.Header.Get("Link")
    if linkHeader == "" {
        return r, nil
    }

    // Split the link header into individual links.
    links := strings.Split(linkHeader, ", ")
    for _, link := range links {
        // Parse the link.
        parts := strings.Split(link, "; ")
        if len(parts) != 2 {
            continue
        }

        // Parse the URL.
        u, err := url.Parse(strings.Trim(parts[0], "<>"
        if err != nil {
            continue
        }

        // Parse the query parameters.
        q := u.Query()
        page := q.Get("page")
        if page == "" {
            continue
        }

        // Parse the page number.
        pageNum, err := strconv.Atoi(page)
        if err != nil {
            continue
        }

        // Determine the relationship.
        rel := strings.Trim(parts[1], "rel=\"\"")
        switch rel {
        case "next":
            r.NextPage = pageNum
        case "prev":
            r.PrevPage = pageNum
        case "first":
            r.FirstPage = pageNum
        case "last":
            r.LastPage = pageNum
        }
    }

    return r, nil
}

// NextPage returns the next page number.
func (r *Response) NextPage() int {
    return r.NextPage
}

// PrevPage returns the previous page number.
func (r *Response) PrevPage() int {
    return r.PrevPage
}

// FirstPage returns the first page number.
func (r *Response) FirstPage() int {
    return r.FirstPage
}

// LastPage returns the last page number.
func (r *Response) LastPage() int {
    return r.LastPage
}
