package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
    values := url.Values{
        "query": {"hello world"},
    }
    resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    log.Println(resp.Status)
    log.Println(resp.StatusCode)
    log.Println(resp.Header)
}

