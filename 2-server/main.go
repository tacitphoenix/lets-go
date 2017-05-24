package main

import (
    "fmt"
    "net/http"
)

func handler (writer http.ResponseWriter, req *http.Request){
    fmt.Fprintf(writer, "Hello world %s", req.URL.Path[1:])
}

func rocks (writer http.ResponseWriter, req *http.Request){
    fmt.Fprintf(writer, "Golang is a System Programming language that rocks!!")
}

func main(){
    http.HandleFunc("/", handler)
    http.HandleFunc("/rocks", rocks)
    http.ListenAndServe(":4040",nil)
}
