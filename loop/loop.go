package main

import (
    "fmt";
    "time"
)

func main() {
    fmt.Printf("Start\n")
    for {
        fmt.Printf("Hello\n")
        time.Sleep(5 * time.Second)
    }
}
