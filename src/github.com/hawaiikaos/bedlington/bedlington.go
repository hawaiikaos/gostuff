package bedlington

import (
    "net/http"
    "html/template"
    //"time"

    //"appengine"
    //"appengine/user"
    //"appengine/datastore"
)

type Page struct {
    Title string
    //Body []byte
}

var templatepath = "templates/"
var templates = template.Must(template.ParseFiles(templatepath+"root.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
    //c := appengine.NewContext(r)
    /*if err := geozoomTemplate.Execute(w, "stuff"); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }*/
    
    var p Page
    p.Title = "Home"
    
    renderTemplate(w, "root", &p)
}

var geozoomTemplate = template.Must(template.New("location").Parse(`
    <html>
    <head>
    <title>Bedlington Geozoom</title>
    </head>
    <body>
    You are here
    </body>
    </html>
`))