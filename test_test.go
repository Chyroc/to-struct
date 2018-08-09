package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func testJSON(t *testing.T, path, expected string) {
	in, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	var buf = new(bytes.Buffer)
	JSONToStruct(buf, "", []byte(in))
	assert.Equal(t, expected, buf.String())
}

func TestJSONToStruct(t *testing.T) {
	testJSON(t, "./testdata/map_1.json", `type Struct struct {
	Code	int
	Count	int
	Result	[]struct {
		Aid	int
		ArtistID	int
		Lrc	string
		Sid	int
		Song	string
	}
}`)

	testJSON(t, "./testdata/slice_1.json", `type Struct struct {
	Slice	[]struct {
		Aid	int
		ArtistID	int
		Lrc	string
		Sid	int
		Song	string
	}
}`)
}

func TestTomlTo(t *testing.T) {
	
}
