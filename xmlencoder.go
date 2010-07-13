package main

import (
  "fmt"
)

type Roar struct {
     Author       string
     Text         string
     CreationDate string
}

func main(){
  var r = &Roar{Author:"Sven", Text:"Hallo, das ist ein Roar!",
  CreationDate:"09.07.2010 09:44"}
  fmt.Printf("Trying to marshal...")
  var result, err = Marshal(r)
  fmt.Printf("done.\n")
  if err != nil {
    fmt.Printf(err.String())
  }
  fmt.Sprintf("%s", string(result))
}
