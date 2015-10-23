package main

import (
    "html/template"
    "io/ioutil"
    "net/http"
    "regexp"
    "fmt"
)

type Page struct {
    Title string
    Body []byte
}

var datapath = "wikiengine/data/"
var templatepath = "wikiengine/templates/"
var templates = template.Must(template.ParseFiles(templatepath+"edit.html", templatepath+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
    filename := datapath + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

/*func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil //title is the second subexpression
}*/

func loadPage(title string) (*Page, error) {
    filename := datapath + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("in viewHandler")
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*func homeHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("in homeHandler")
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
    http.Redirect(w, r, "/view/Frontpage", http.StatusFound)
}*/
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in homeHandler")
    fmt.Fprintf(w, "<h1>stuff<h1>")
    //http.Redirect(w, r)
}

func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    http.HandleFunc("/stuff/", homeHandler)
    http.ListenAndServe(":9072", nil)
}

