package main

import (
	"github.com/BurntSushi/toml"

	"bytes"
)

func tomlTo(buf *bytes.Buffer, padding string, in []byte) {
	var m = make(map[string]interface{})
	toml.Decode(string(in), &m)

	mapToStruct(buf, padding, m)
}
