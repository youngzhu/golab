package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// https://golang.google.cn/doc/articles/wiki/

type Page struct {
	Title string
	Body  []byte
}

func (p Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	//fmt.Println(os.Getwd())
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// URL: http://localhost:8080/monkeys
	// Output: Hi there, I love monkeys!
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func _viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	renderTemplate(w, "view", p)
}

// hard-coded
func _editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	format := `
		<h1>Editing %s</h1>
		<form action="/save/%s" method="POST">
			<textarea name="body">%s</textarea><br>
			<input type="submit" value="Save">
		</form>
	`
	fmt.Fprintf(w, format, p.Title, p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("wiki/" + tmpl + ".html")
	if err != nil {
		fmt.Println("template error:", err)
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
