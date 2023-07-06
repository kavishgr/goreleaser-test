package main

import(
  "fmt"
  "runtime"
)

func main(){
  os := runtime.GOOS
  fmt.Println("Hello from Go on", os)
  fmt.Println("Goodbye!")
}
