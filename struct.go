package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/serenize/snaker"
)

func mapToStruct(buf *bytes.Buffer, padding string, m map[string]interface{}) {
	if padding == "" {
		buf.WriteString("type Struct ")
	}
	buf.WriteString("struct {\n")
	newPadding := padding + "\t"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		n := new(bytes.Buffer)
		interfaceToType(n, newPadding, m[k])

		buf.WriteString(fmt.Sprintf(newPadding+"%s\t%s\n", snaker.SnakeToCamel(k), n.String()))
	}
	buf.WriteString(padding + "}")
}

func sliceToStruct(buf *bytes.Buffer, padding string, m []interface{}) {
	if padding == "" {
		buf.WriteString("type Struct ")
	}
	buf.WriteString("struct {\n")
	newPadding := padding + "\t"

	n := new(bytes.Buffer)
	interfaceToType(n, newPadding, m[0])
	buf.WriteString(fmt.Sprintf(newPadding+"%s\t[]%s\n", "Slice", n.String()))
	buf.WriteString(padding + "}")
}
