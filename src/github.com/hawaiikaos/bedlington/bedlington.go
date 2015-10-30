package bedlington

import (
    "net/http"
    "html/template"
    "fmt"
    //"time"

    "appengine"
    //"appengine/user"
    "appengine/datastore"
)

type Page struct {
    Title string
    //Body []byte
}

type Topic struct {
    TopicName string
}

type Vote struct {
    Topic Topic
    Cast string
}

var templatepath = "templates/"
var templates = template.Must(template.ParseFiles(  templatepath+"root.html",
                                                    templatepath+"vote.html",
                                                    templatepath+"addtopic.html",
                                                    templatepath+"tally.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/vote/", vote)
    http.HandleFunc("/vote/addtopic/", addtopic)
    http.HandleFunc("/vote/tally/", tally)
    http.HandleFunc("/test/", test)
}

func vote(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in vote function")
    var p Page
    p.Title = "Vote"
    renderTemplate(w, "vote", &p)
}

func test(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in test")
    fmt.Fprintf(w, "test")
    
}

func topicKey(c appengine.Context) *datastore.Key {
    return datastore.NewKey(c, "Topics", "", 0, nil)
}

func voteKey(c appengine.Context) *datastore.Key {
    return datastore.NewKey(c, "Votes", "", 0, nil)
}

func addtopic(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("in add topic")
    var p Page
    p.Title = "Add Topic"
    
    c := appengine.NewContext(r)
    q := datastore.NewQuery("Topic").Ancestor(topicKey(c)).Order("-Date").Limit(20)
    topics := make([]Topic, 0, 20)
    if _, err := q.GetAll(c, &topics); err != nil {
        fmt.Println("in error")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } else {
        renderTemplate(w, "addtopic", &p)
    }
    /*if err := renderTemplate(w, "addtopic", &p) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }*/
}

func tally(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in tally")
    var p Page
    p.Title = "Tally"
    renderTemplate(w, "tally", &p)
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in root function")
    //c := appengine.NewContext(r)
    /*if err := geozoomTemplate.Execute(w, "stuff"); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }*/
    
    var p Page
    p.Title = "Home"
    
    renderTemplate(w, "root", &p)
}