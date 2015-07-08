package main

import (
  "fmt"
)

var c = make(chan int, 1) // 1
var a string

func f() {
        a = "hello world tgo"
        <-c
}

func main() {
        go f()
        c <- 0 // 2
        c <- 3  // 3
        fmt.Println(a)
}
