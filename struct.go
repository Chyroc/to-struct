package main

import (
	"sort"
	"fmt"
	"strings"
	"bytes"
)

func mapToStruct(buf *bytes.Buffer, padding string, m map[string]interface{}) {
	if padding == "" {
		buf.WriteString("type Struct struct {\n")
	} else {
		buf.WriteString("struct {\n")
	}
	newPadding := padding + "\t"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.StringsAreSorted(keys)

	for _, k := range keys {
		n := new(bytes.Buffer)
		interfaceToType(n, newPadding, m[k])
		buf.WriteString(fmt.Sprintf(newPadding+"%s\t%s\n", strings.Title(k), n.String()))
	}
	buf.WriteString(padding + "}")
}
