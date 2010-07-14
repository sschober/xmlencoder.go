# XML Encoder for Go

An XML encoder for go. Heavily inspired by the json encoder that
comes with the go distribution.

## Usage

From `encoder_test.go`:

    type Roar struct {
        Author       string
        Text         string
        CreationDate string
    }

    func TestMarshal(t *testing.T) {
        var r = &Roar{Author: "Sven", Text: "Hallo, das ist ein Roar!",
            CreationDate: "09.07.2010 09:44"}
        result, err := Marshal(r)
        if err != nil {
            t.Fatalf("Marshal r: %v", err)
        }
        fmt.Printf("%s", string(result))
    }

yields:

    <Roar><Author>Sven</Author><Text>Hallo, das ist ein Roar!</Text><CreationDate>09.07.2010 09:44</CreationDate></Roar>

## Note

I'm aware of Russ Coxs' [Statement][mail]:

> XML is for formatting documents, not for describing data
> structures.  The fact that it gets used often for the latter
> doesn't lessen its fundamental inability to do the job well.

And i think the omission of a `Marshal()` function is no
coincidence, but ...

## Author

Sven Schober

[mail]: http://groups.google.com/group/golang-nuts/msg/3a1efaf073cc692c?
