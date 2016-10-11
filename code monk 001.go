package main

import (
	"fmt"
  //"strings"
)

  func main() {
  var n int

  fmt.Scanln(&n)
  for n > 0{
      var s string
      fmt.Scanln(&s)
      size_s := len(s)
      for i := 0; i <len(s); i++{
        size_s--
        index_s:= string([]rune(s)[size_s])
    		fmt.Print(index_s)
    	}
      fmt.Println()
      n--
  }
}
