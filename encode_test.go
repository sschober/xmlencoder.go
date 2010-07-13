package xmlencoder

import (
    "testing"
    "fmt"
)

type Roar struct {
    Author       string
    Text         string
    CreationDate string
}

func TestMarshal(t *testing.T) {
    var r = &Roar{Author: "Sven", Text: "Hallo, das ist ein Roar!",
        CreationDate: "09.07.2010 09:44"}
    fmt.Printf("Trying to marshal...")
    result, err := Marshal(r)
    fmt.Printf("done.\n")
    if err != nil {
        t.Fatalf("Marshal allValue: %v", err)
    }
    fmt.Printf("%s", string(result))
}
