package main

import(
  "fmt"
  "runtime"
)

func main(){
  os := runtime.GOOS
  fmt.Println("Hello from", os)
}
