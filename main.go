package main

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"bytes"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var buf = new(bytes.Buffer)

	JSONToStruct(buf, "", in)

	fmt.Println(buf.String())
}
