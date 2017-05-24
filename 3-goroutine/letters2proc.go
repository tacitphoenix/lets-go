package main

import(
    "fmt"
    "runtime"
    "sync"
)

func main(){
    // allocate 1 logical processor for the scheduler
    logical_processors := 2
    runtime.GOMAXPROCS(logical_processors)
    num := 3

    // wg wait for added goroutines to finish
    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Start Goroutines")

    // declare goroutine for small letters
    go func(){
        // tell wait group when done
        defer wg.Done()

        // display alphabet three times
        for n := 0; n < num; n++ {
          for char := 'a'; char < 'a'+26; char++{
              fmt.Printf("%c ", char)
          }
        }
    }()

    // declare goroutine for large letters
    go func(){
        // tell wait group when done
        defer wg.Done()

        // display alphabet three times
        for n := 0; n < num; n++ {
          for char := 'A'; char < 'A'+26; char++{
              fmt.Printf("%c ", char)
          }
        }
    }()

    // waiting to finish
    fmt.Println("Waiting To Finish")
    wg.Wait()
    fmt.Println("\nTerminating Program")
}
