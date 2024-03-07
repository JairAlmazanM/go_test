package main

import (
    "html/template"
    "net/http"
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

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}