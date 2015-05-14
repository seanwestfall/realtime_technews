package main

// Title of SPA
const spaTitle string = "RealTime Tech News"

import (
    "fmt"           // standard
    "io/ioutil"     // io lib
    "html/template" // allows for html templates to be imported
    "net/http"      // the server 

    "golang.org/x/net/websocket" // the websocket library
)

// The data structure that holds a page
type Page struct {
    Title string // title of the page
    Body []byte // HTML body
}

// The function that loads an html page
func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, _err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Boby: body}, nil
}

// This is the main index page
func viewIndex( w http.ResponseWritter, r *http.Request) {
    title := spaTitle
}

// Application entry
func main() {
    http.HandleFunc("/", viewIndex)
    http.ListenAndServe(":8080", nil) // listen on port 8080
}
