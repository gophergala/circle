package main

import (
    "testing"
)

//func TestSort(t *testing.T){
//    err := Sort("/home/mohan/Desktop/httpservercef.go")
//    
//    if err == nil {
//        t.Fatal("File gets recognized as directory")
//    }
//    
//    err = Sort("/home/mohan/testbed")
//    
//    if err != nil {
//        t.Fatal("Directory doesn't get recognized")
//    }
//}

func TestSortWithRegexp(t *testing.T){
    err := SortWithRegexp("/home/mohan/Desktop/", "^s");
    
    if err != nil {
        t.Fatal("sort failed")
    }
}

//func TestCreate(t *testing.T){
//    CreateINE("/home/mohan/Desktop/mega")
//}