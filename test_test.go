package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func testToStruct(t *testing.T, path, expected string) {
	in, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, expected, toStruct(in))
}

func TestJSONToStruct(t *testing.T) {
	testToStruct(t, "./testdata/map_1.json", `type Struct struct {
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

	testToStruct(t, "./testdata/slice_1.json", `type Struct struct {
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
	testToStruct(t, "./testdata/toml_1.toml", `type Struct struct {
	Clients	struct {
		Data	[][]string
		Hosts	[]string
	}
	Database	struct {
		ConnectionMax	int64
		Enabled	bool
		Ports	[]int64
		Server	string
	}
	Owner	struct {
		Dob	time.Time
		Name	string
	}
	Servers	struct {
		Alpha	struct {
			Dc	string
			IP	string
		}
		Beta	struct {
			Dc	string
			IP	string
		}
	}
	Title	string
}`)
}
