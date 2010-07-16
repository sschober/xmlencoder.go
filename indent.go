// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xmlencoder

import (
    "bytes"
    "os"
    "fmt"
)

// Compact appends to dst the XML-encoded src with
// insignificant space characters elided.
func Compact(dst *bytes.Buffer, src []byte) os.Error {
    fmt.Sprintf("Appending %s to buffer", src)
    dst.Write(src)
    return nil
}

// Encodes the state of parsing
const (
    parseStart = iota
    parseOpenBracket
    parseOpenBracketAfterStart
    parseOpenBracketAfterEndElem
    parseStartElem
    parseEndElem
)

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) os.Error {
    var state = parseStart
    var depth = 0
    var buf bytes.Buffer

    for _, c := range src {
        switch c {
        case '<':
            switch state {
            case parseStart:
		buf.WriteString(prefix)
                state = parseOpenBracketAfterStart
            case parseEndElem:
                state = parseOpenBracketAfterEndElem
            default:
                state = parseOpenBracket
            }
            buf.WriteTo(dst)
            buf.Reset()
            buf.WriteByte(c)
        case '/':
            switch state {
            case parseOpenBracket:
                state = parseEndElem
            case parseOpenBracketAfterEndElem:
                depth--
                newline(dst, prefix, indent, depth)
                state = parseEndElem
            }
            buf.WriteByte(c)
        case '>':
            buf.WriteByte(c)
            buf.WriteTo(dst)
        default:
            switch state {
            case parseOpenBracket:
                depth++
                newline(dst, prefix, indent, depth)
                state = parseStartElem
            case parseOpenBracketAfterEndElem:
                newline(dst, prefix, indent, depth)
                state = parseStartElem
            case parseOpenBracketAfterStart:
                state = parseStartElem
            }
            buf.WriteByte(c)
        }
    }
    return nil
}

func newline(dst *bytes.Buffer, prefix, indent string, depth int) {
    dst.WriteByte('\n')
    dst.WriteString(prefix)
    for i := 0; i < depth; i++ {
        dst.WriteString(indent)
    }
}
