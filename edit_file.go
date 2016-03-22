package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func dirReader(dir string) string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    dirElements := make(map[string]int64)
    for _, file := range files {
        dirElements[file.Name()] = file.Size()
    }
    return dirElements
}

func main() {
    fmt.Println(dirReader("."))
}