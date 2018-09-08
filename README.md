# protoc-go-tagger

[![Build Status](https://travis-ci.org/rezam90/protoc-go-tagger.svg?branch=master)](https://travis-ci.org/rezam90/protoc-go-tagger)
[![Go Report Card](https://goreportcard.com/badge/github.com/rezam90/protoc-go-tagger)](https://goreportcard.com/report/github.com/rezam90/protoc-go-tagger)
[![Coverage Status](https://coveralls.io/repos/github/rezam90/protoc-go-tagger/badge.svg)](https://coveralls.io/github/rezam90/protoc-go-tagger)

## Why?

Golang [protobuf](https://github.com/golang/protobuf) doesn't support
[custom tags to generated structs](https://github.com/golang/protobuf/issues/52). This
script injects custom tags to generated protobuf files, useful for
things like validation struct tags.

## Install

* [protobuf version 3](https://github.com/google/protobuf)

  For OS X:
  
  ```
  brew install protobuf
  ```
* go support for protobuf: `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

*  `go get github.com/rezam90/protoc-go-tagger` or download the
  binaries from releases page.

## Usage

Add a comment with syntax `// @tag: custom_tag:"custom_value"`
before fields to add custom tag to in .proto files.

Example:

```
// file: test.proto
syntax = "proto3";

package pb;

message IP {
  // @tag: valid:"ip"
  string Address = 1;
}
```

Generate with protoc command as normal.

```
protoc --go_out=. test.proto
```

Run `protoc-go-tagger` with generated file `test.pb.go`.

```
protoc-go-tagger -input=./test.pb.go
```

The custom tags will be injected to `test.pb.go`.

```
type IP struct {
	// @tag: valid:"ip"
	Address string `protobuf:"bytes,1,opt,name=Address,json=address" json:"Address,omitempty" valid:"ip"`
}
```

To skip the tag for the generated XXX_* fields, use
`-XXX_skip=yaml,xml` flag.
