package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page structure represents pages of the wiki
type Page struct {
	Title string
	Body  []byte // byte is the type expected by the io libraries
}

// save method on page for data percistancy
// The save method returns an error value because that is the return type of WriteFile
// If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types).
// func (p *Page) saves pages into a text file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // write content of p.Body into the file called 'filename'
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

// func viewHandler will handle URLs prefixed with "/view/".
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title) // _ ignores error return from loadPage,
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler) // must have the file hello.txt to be loaded and served
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// WHY THE FILE.TXT NEEDS TO HAVE THE SAME NAME AS THE URL PATH???
// https://golang.org/doc/articles/wiki/  (Editing Pages) - @here
