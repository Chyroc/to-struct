package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func toStruct(in []byte) string {
	var buf = new(bytes.Buffer)

	// map slice
	JSONToStruct(buf, "", in)
	if s := buf.String(); s != "" {
		return s
	}

	// toml
	tomlTo(buf, "", in)
	if s := buf.String(); s != "" {
		return s
	}

	return ""
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toStruct(in))
}
