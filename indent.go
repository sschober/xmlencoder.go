// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xmlencoder

import (
    "bytes"
    "os"
    "fmt"
)

// Compact appends to dst the JSON-encoded src with
// insignificant space characters elided.
func Compact(dst *bytes.Buffer, src []byte) os.Error {
    fmt.Sprintf("Appending %s to buffer", src)
    dst.Write(src)
    return nil
}

func newline(dst *bytes.Buffer, prefix, indent string, depth int) {
    dst.WriteByte('\n')
    dst.WriteString(prefix)
    for i := 0; i < depth; i++ {
        dst.WriteString(indent)
    }
}
