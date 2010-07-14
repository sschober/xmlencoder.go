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

func NewRoar(author, text, creationDate string) *Roar {
  return &Roar{author, text, creationDate}
}

var r = NewRoar("Sven", "Bruell!", "14.07.2010 07:31")

func TestMarshal(t *testing.T) {
    result, err := Marshal(r)
    if err != nil {
        t.Fatalf("Marshal r: %v", err)
    }
    fmt.Printf("%s\n", string(result))
}

func TestMarshalAndNamespace(t *testing.T) {
    mapping := NewNSMap(r, "myns")
    result, err := MarshalWithNSMap(r, mapping)
    if err != nil {
        t.Fatalf("Marshal r: %v", err)
    }
    fmt.Printf("%s\n", string(result))
}
