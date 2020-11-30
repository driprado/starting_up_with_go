package main
 import (
	 "fmt"
	 "io/ioutil"
 )

 // https://golang.org/doc/articles/wiki/  (Callers of this function) - @here


 // This structure represents pages of the wiki
 type Page struct {
	 Title string
	 Body []byte // byte is the type expected by the io libraries
 }

 // save method on page for data percistancy
 // The save method returns an error value because that is the return type of WriteFile 
 // If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types).
 // func (p *Page) saves pages into a text file
 func (p *Page) save() error {
	 filename := p.Title + ".txt"
	 return ioutil.WriteFile(filename, p.Body, 0600)
 }

 //func loadPage instanciates a page from filename
 // returns an instance of Page or error
 func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
 }

