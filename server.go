package main

import (
    "io"
    "os"
    "log"
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "github.com/mitchellh/go-homedir"
    
)

type FList struct {
    Files []os.FileInfo
}
    

func Run(){
    
    
    r := mux.NewRouter()
    
    r.HandleFunc("/", DefaultHandler).Methods("GET")
    r.HandleFunc("/", PathHandler).Methods("POST")
    
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

func DefaultHandler(res http.ResponseWriter, req *http.Request){
    path, err := homedir.Dir()
    
    if err != nil {
        log.Fatal(err)
    }
    
    dir, err := os.Open(path)
    
    if err != nil {
        log.Fatal(err)
    }
    
    fi, err := dir.Readdir(100)
    
    if err != nil {
        log.Fatal(err)
    }
    
    
    
    t, err := template.ParseFiles("index.html")
    if err != nil {
        log.Fatal(err)
    }
    
    fobj := &FList{Files: fi}
    
    err = t.Execute(res, fobj)
    if err != nil {
        log.Fatal(err)
    }
}

func PathHandler(res http.ResponseWriter, req *http.Request){
    newpath := req.FormValue("path")
    w := io.Writer(res)
    io.WriteString(w, newpath)
}