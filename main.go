package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// slice the part with the view, and just return the stuff that comes after view
	title := r.URL.Path[len("/view/"):] // this extracts the title from the request URL
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Body)
}

// this is a method named save, that takes at it´s reveiver p,
// a pointer to Page. It takes no Parameters and returns a value of type error
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
	// this writes a file with given filename, the byte slice from page, and read+write rights
	// for the current user
}

// func loadPage constructs the file name from the title Parameter
// reads the files content into a new variable body and returns a pointer
// to Page
// if the caller of this function gets nil as an return value
// the page is loaded correctly, if not the caller can handle the error
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
