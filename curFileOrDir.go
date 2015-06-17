package main

import (
  "fmt"
  "runtime"
  "path"
)

// __FILE__
func GetCurrentFile() string {
  _, filename, _, _ := runtime.Caller(1)
  return filename
}

// __DIR__
func GetCurrentDir() string {
  _, filename, _, _ := runtime.Caller(1)
  return path.Dir(filename)
}

func main() {
        fmt.Println("current file", GetCurrentFile())
        fmt.Println("current dir", GetCurrentDir())
        
        _, filename, _, _ := runtime.Caller(1)
        fmt.Println(filename)
}
