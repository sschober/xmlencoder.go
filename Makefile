# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=xmlencoder
GOFILES=\
	indent.go\
	encode.go

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

include $(GOROOT)/src/Make.pkg

fmt:
	for src in *.go; do \
		${GOFMT} -w $$src; \
	done
