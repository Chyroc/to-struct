package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var buf = new(bytes.Buffer)

	// map slice
	JSONToStruct(buf, "", in)
	if s := buf.String(); s != "" {
		fmt.Println(s)
	}

}
