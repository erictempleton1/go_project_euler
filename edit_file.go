package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func dirReader(dir string) []string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    var s []string
    for _, file := range files {
        s = append(s, file.Name())
    }
    return s
}

func main() {
    fmt.Println(dirReader("."))
}