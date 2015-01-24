package main

import (
    "testing"
)

func TestSort(t *testing.T){
    err := Sort("/home/mohan/Desktop/httpservercef.go")
    
    if err == nil {
        t.Fatal("File gets recognized as directory")
    }
    
    err = Sort("/home/mohan/Desktop")
    
    if err != nil {
        t.Fatal("Directory doesn't get recognized")
    }
}