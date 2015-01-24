package main

import (
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
    r.HandleFunc("/sort", SortHandler).Methods("POST")
    
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

func DefaultHandler(res http.ResponseWriter, req *http.Request){
    path, err := homedir.Dir()
    log.Println(path)
    checkErr(err)
    
    fobj := CreateFList(path)
    
    t, err := template.ParseFiles("index.html")
    checkErr(err)
    
    err = t.Execute(res, fobj)
    checkErr(err)
}

func PathHandler(res http.ResponseWriter, req *http.Request){
    path := req.FormValue("path")
    log.Println(path)
    fobj := CreateFList(path)
    
    t, err := template.ParseFiles("index.html")
    checkErr(err)
    
    err = t.Execute(res, fobj)
    checkErr(err)
}

func SortHandler(res http.ResponseWriter, req *http.Request){
    path := req.FormValue("path")
    log.Println(path)
    Sort(path)
}

func checkErr(err error){
    if err != nil {
        log.Fatal(err)
    }
}

func CreateFList(path string) *FList {
    dir, err := os.Open(path)
    checkErr(err)
    
    fi, err := dir.Readdir(100)
    checkErr(err) 
    
    fobj := &FList{Files: fi}
    
    return fobj
}
    
    