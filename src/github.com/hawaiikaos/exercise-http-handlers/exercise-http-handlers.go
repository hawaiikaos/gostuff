package main

import (
    //"log"
    "net/http"
    "fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
        fmt.Fprintf(w, s.Greeting)
        fmt.Fprintf(w, s.Punct)
        fmt.Fprintf(w, s.Who)
        //fmt.Println(s.Who)
    }

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
        fmt.Fprintf(w, "%s", s)
        //fmt.Println(s)
    }

func main() {
    //your http.Handle calls here
    
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}

/*func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4005", h)
	if err != nil {
		log.Fatal(err)
	}
}*/