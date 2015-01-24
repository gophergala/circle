package main

import (
    "os"
    "log"
    "errors"
)

func checkErr(err error){
    if err != nil {
        log.Println(err)
    }
}

func Sort(path string) error {
    handle, err := os.Lstat(path)
    
    if !handle.IsDir() {
        log.Print("It's not a directory")
        return errors.New("error")
    }
    
    dir, err := os.Open(path)
    fi, err := dir.Readdir(100)
    for _, file := range fi {
        log.Println(file.Name())
    }
        
    
    return err
    
}
    
    
    
    
    