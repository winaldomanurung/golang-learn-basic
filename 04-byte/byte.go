package main

import "fmt"

func main() {
//byte type in Golang is an alias for the unsigned integer 8 type (uint8)
//The range of a byte is 0 to 255 (same as uint8)

  var a byte = 100
  var b byte = 'B'
  var c byte = 'C'

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}