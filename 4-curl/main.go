package main

import(
    "io"
    "log"
    "net/http"
    "os"
)

func main(){
    // get response
    r, err := http.Get(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    // create file
    file, err := os.Create(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    // write to both stdout and file outputs
    dest := io.MultiWriter(os.Stdout, file)

    // read response and send to both outputs
    io.Copy(dest, r.Body)
    if err := r.Body.Close(); err != nil {
        log.Println(err)
    }

}
