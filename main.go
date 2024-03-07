package main

import (
    "html/template"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.HandleFunc("/path/to/content", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is the content loaded by htmx."))
    })

    http.Handle("/static/", correctContentType(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func correctContentType(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, ".css") {
            w.Header().Add("Content-Type", "text/css")
        }
        next.ServeHTTP(w, r)
    })
}