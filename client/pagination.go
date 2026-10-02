package client

import (
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

type ListOptions struct {
    Page    int `url:"page,omitempty"`
    PerPage int `url:"per_page,omitempty"`
}

type Response struct {
    *http.Response
    NextPage   int
    PrevPage   int
    FirstPage  int
    LastPage   int
}

func parsePagination(resp *http.Response) (*Response, error) {
    r := &Response{Response: resp}
    linkHeader := resp.Header.Get("Link")
    if linkHeader == "" {
        return r, nil
    }
    
    links := strings.Split(linkHeader, ",")
    for _, link := range links {
        parts := strings.Split(link, ";")
        if len(parts) < 2 {
            continue
        }
        
        urlPart := strings.Trim(parts[0], " <>\"")
        relPart := strings.Trim(parts[1], " ")
        
        if !strings.HasPrefix(relPart, "rel=\"") {
            continue
        }
        
        rel := strings.Trim(relPart[5:], "\"")
        u, err := url.Parse(urlPart)
        if err != nil {
            continue
        }
        
        pageStr := u.Query().Get("page")
        if pageStr == "" {
            continue
        }
        
        page, err := strconv.Atoi(pageStr)
        if err != nil {
            continue
        }
        
        switch rel {
        case "next":
            r.NextPage = page
        case "prev":
            r.PrevPage = page
        case "first":
            r.FirstPage = page
        case "last":
            r.LastPage = page
        }
    }
    
    return r, nil
}

func (r *Response) NextPage() int   { return r.NextPage }
func (r *Response) PrevPage() int   { return r.PrevPage }
func (r *Response) FirstPage() int  { return r.FirstPage }
func (r *Response) LastPage() int   { return r.LastPage }
