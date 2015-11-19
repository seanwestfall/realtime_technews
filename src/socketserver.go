/**
 * This Source Code is licensed under the MIT license. If a copy of the
 * MIT-license was not distributed with this file, You can obtain one at:
 * http://opensource.org/licenses/mit-license.html.
 *
 * @author: Sean Westfall
 * @license MIT
 * @copyright Sean Westfall, 2015
 * 
 * name: RealTime Tech News
 * version: 0.0.0
 * description: get tech news in real time
 * homepage: fieldsofgoldfish.com
 */
package main

import (
  "flag"			// command parser
  "fmt"           // standard
  "io/ioutil"     // io lib
  "html/template" // allows for html templates to be imported
	"net"
  "net/http"      // the server 

	"log"			// log errors

  //"golang.org/x/net/websocket" // the websocket library
)

// Title of SPA
const spaTitle string = "RealTime Tech News"
const indexFile string = "index.html"
const contentFile string = "lib/src/main.tpl"

var (
  addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)


// The data structure that holds a page
type Page struct {
    Title string // title of the page
    Body []byte // body in bytes
    Template template.HTML // body in html
}

// The function that loads an html page
func loadPage(filename string, title string) (*Page, error) {
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body, Template: template.HTML(body)}, nil
}

// template cache
// (add to this list to cache)
var templates = template.Must(template.ParseFiles(indexFile))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// check to see that the path exists using a regex, (replace validpaths)
//var validPath = regexp.MustCompile("^/validpaths/([a-zA-Z0-9]+)$")

// a generic handler maker (TODO: needs work)
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//m := validPath.FindStringSubmatch(r.URL.Path)
    //if m == nil {
    //	http.NotFound(w, r)
    //	return
    //}
    //fn(w, r, m[2])

		// this is a temp comment out until the generic handler is developed further
		fn(w, r)
	}
}

/*
 * This is the main index page
 * the content and title are defined in the const definitions above
 * - contentFile
 * - spaTitle
 * - indexFile
 */
func viewIndex(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(contentFile, spaTitle) // load the main page and set the title
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
        return
	}
	renderTemplate(w, indexFile, p)
}

// Application entry
func main() {
	flag.Parse() // parse the cli

	// router and generate handlers
    http.HandleFunc("/", makeHandler(viewIndex))

	if *addr {	
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}
	
  // start the http server
  fmt.Printf("Server to PORT: 8080\n")
  http.ListenAndServe(":8080", nil) // listen on port 8080
}
